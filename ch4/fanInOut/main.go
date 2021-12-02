package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 扇入扇出
func main() {
	toInt:= func(done <-chan interface{},valueStream <-chan interface{})<-chan int {
		intStream:=make(chan int)
		go func() {
			defer close(intStream)
			for v:=range valueStream{
				select {
				case <-done:
					return
				case intStream<-v.(int):
				}
			}
		}()
		return intStream
	}
	repeatFn:= func(done <-chan interface{},fn func()interface{})<-chan interface{} {
	  valueStream:=make(chan interface{})
	  go func() {
	  	defer close(valueStream)
	  	for{
			select {
			case <-done:
				return
			case valueStream<-fn():
			}
		}
	  }()
	  return valueStream
	}
	take:= func(done <-chan interface{},valueStream <-chan interface{},num int)<-chan interface{}{
		takeStream:=make(chan interface{})
		go func() {
			defer close(takeStream)
			for i:=0;i<num;i++{
				select {
				case <-done:
					return
				case takeStream<-<-valueStream:
				}
			}
		}()
		return takeStream
	}
	primeFinder:= func(done <-chan interface{},randIntStream <-chan int)<-chan interface{} {
		valStream:=make(chan interface{})
		go func() {
			defer close(valStream)
			for i:=range randIntStream{
				select {
				case <-done:
					return
				case valStream<-i:
				}
			}
		}()
		return valStream
	}
	fanIn:= func(done <- chan interface{},channels ...<-chan interface{})<-chan interface{} {  // 扇入方式，汇总channel
		var wg sync.WaitGroup
		multiplexedStream:=make(chan interface{})
		multiplex:= func(c <-chan interface{}) {
			defer wg.Done()
			for i:=range c{
				select {
				case <-done:
					return
				case multiplexedStream<-i:
				}
			}
		}
		// 从所有的通道中取数据
		wg.Add(len(channels))
		for _,c:=range channels{
			go multiplex(c)
		}
		// 等待所有数据汇总完毕
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream

	}
    rands:= func()interface{}{return rand.Intn(50000000)}
     done:=make(chan interface{})
     defer close(done)
     start:=time.Now()
     randIntStream:=toInt(done,repeatFn(done,rands))
     //fmt.Println("Primes:")
     //for prime:=range take(done,primeFinder(done,randIntStream),10){
     //	fmt.Printf("\t%d\n",prime)
	 //}
     //fmt.Printf("search took : %v \n",time.Since(start))

     numFinders:=runtime.NumCPU()
     fmt.Printf("Spinning up %d prime finders.\n",numFinders)
     finders:=make([]<-chan interface{},numFinders)
     fmt.Println("Primes:")
     for i:=0;i<numFinders;i++{
     	finders[i]=primeFinder(done,randIntStream)
	 }
     for prime:=range take(done,fanIn(done,finders...),10){
     	fmt.Printf("\t%d \n",prime)
	 }
     fmt.Printf("Search took : %v ",time.Since(start))

}

