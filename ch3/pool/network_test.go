package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

// go test -benchtime=10s -bench=.

//在使用Pool时，请记住以下几点：
//
//实例化sync.Pool时，给它一个新元素，该元素应该是线程安全的。
//当你从Get获得一个实例时，不要假设你接收到的对象状态。
//当你从池中取得实例时，请务必不要忘记调用Put。否则池的优越性就体现不出来了。这通常用defer来执行延迟操作。
//池中的元素必须大致上是均匀的。

func connectToService() interface{}  {
	time.Sleep(1*time.Second)
	return struct {}{}
}

func StartWorkDaemon() *sync.WaitGroup  {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server,err:=net.Listen("tcp","localhost:8081")
		if err !=nil{
			log.Fatalf("cannot listen :%v ",err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn,err:=server.Accept()
			if err !=nil{
				log.Printf("cannot accept connection: %v",err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn,"")
			conn.Close()
		}

	}()
  return &wg
}

func init()  {
	daemonStarted:=StartWorkDaemon()
	daemonStarted.Wait()
	
	daemonCacheStarted:=StartNetworkDaemonCache()
	daemonCacheStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B)  {
	for i:=0;i<b.N;i++{
		conn,err:=net.Dial("tcp","localhost:8081")
		if err !=nil{
			b.Fatalf("cannot dial host : %v",err)
		}
		if _,err:=ioutil.ReadAll(conn);err!=nil{
			b.Fatalf("cannot read : %v",err)
		}
		conn.Close()
	}
}

func warmServiceConnCache() *sync.Pool {
	p:=&sync.Pool{
		New:connectToService,
	}
	for i:=0;i<10;i++{
		p.Put(p.New())
	}
	return p
}

func StartNetworkDaemonCache()*sync.WaitGroup  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool:=warmServiceConnCache()
		server,err:=net.Listen("tcp","localhost:8082")
		if err !=nil{
			log.Fatalf("cannnot listen : %v",err)
		}
        defer server.Close()
		wg.Done()
		for {
			conn,err:=server.Accept()
			if err !=nil{

			}
			svcConn:=connPool.Get()
			fmt.Fprintln(conn,"")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func BenchmarkNetworkRequestCache(b *testing.B)  {
	for i:=0;i<b.N;i++{
		conn,err:=net.Dial("tcp","localhost:8082")
		if err !=nil{
			b.Fatalf("cannot dial host : %v",err)
		}
		if _,err:=ioutil.ReadAll(conn);err!=nil{
			b.Fatalf("cannot read : %v",err)
		}
		conn.Close()
	}
}