package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 请求并发复制处理
	doWork:= func(done <-chan interface{},id int,wg *sync.WaitGroup,result chan<- int) {
          started :=time.Now()
          defer wg.Done()

          // 模拟随机加载
          simulatedLoadTime:=time.Duration(1+rand.Intn(5))*time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		  }

		select {
		case <-done:
		case result<-id:
		  }

          took:=time.Since(started)
          if took<simulatedLoadTime{
          	took=simulatedLoadTime
		  }
          fmt.Printf("%v took %v\n",id,took)
	}

	done:=make(chan interface{})
	result:=make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go doWork(done,i,&wg,result)
	}

	firstReturned:=<-result // 抓取处理程序第一个返回的值。
	close(done) // 取消所有剩余的处理程序。这确保他们不会继续做不必要的工作。
	wg.Wait()
	fmt.Printf("Received an answer from #%v\n",firstReturned)


}

