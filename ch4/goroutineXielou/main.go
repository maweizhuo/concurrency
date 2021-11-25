package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 输出Done. doWork 永不会执行到
	//doWork:= func(strings <-chan string)<-chan interface{} {
    //    completed:=make(chan interface{})
    //    go func() {
    //    	defer fmt.Println("doWork exited")
    //    	defer close(completed)
    //    	for s:=range strings{
    //    		fmt.Println(s)
	//		}
	//	}()
    //    return completed
	//}
	//doWork(nil)
	//// 这里还有其他任务执行
	//fmt.Println("Done.")

	// 尽管向doWork传递了nil给strings通道，我们的goroutine依然正常运行至结束。与之前的例子不同，本例中我们把两个goroutine连接在一起之前，我们建立了第三个goroutine以取消doWork中的goroutine，并成功消除了泄漏问题。
	//doWork:= func(done <-chan interface{},strings <-chan string)<-chan interface{}{
	//	terminated:=make(chan interface{})
	//	go func() {
	//		defer fmt.Println("doWork exited.")
	//		defer close(terminated)
	//		for {
	//			select {
	//			 case s:=<-strings:
	//			 	// do something interesting
	//			 	fmt.Println(s)
	//			case <-done:
	//				return
	//			}
	//		}
	//	}()
	//	return terminated
	//}
	//done:=make(chan interface{})
	//terminated:=doWork(done,nil)
	//go func() {
	//	// cancel the operation after 1 second
	//	time.Sleep(1*time.Second)
	//	fmt.Println("Canceling doWork goroutine...")
	//	close(done)
	//}()
    //<-terminated
    //fmt.Println("Done.")

    // 向通道写入时阻塞goroutine会怎样
    //newRandStream:= func()<-chan int{
    //	randStream:=make(chan int)
    //	go func() {
    //		defer fmt.Println("newRandStream closure exited.") // 该条语句未执行，因为没有告诉他停下来
    //		defer close(randStream)
    //		for {
    //			randStream<-rand.Int()
	//		}
	//	}()
    //	return randStream
	//}
	//
    //randStream:=newRandStream()
    //fmt.Println("3 random ints :")
    //for i:=1;i<=3;i++{
    //	fmt.Printf("%d: %d\n",i,<-randStream)
	//}

	// 上面的解决办法: 为生产者提供一条通知他退出的通道
	newRandStream:= func(done <-chan interface{})<-chan int {
		randStream:=make(chan int)
		go func() {
			defer fmt.Println("newRandStream Closure exited.")
			defer close(randStream)
			for  {
				select {
				case randStream<-rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}
	done:=make(chan interface{})
	randStream:=newRandStream(done)
	fmt.Println("3 random ints :")
	for i:=1;i<=3;i++{
		fmt.Printf("%d: %d \n",i,<-randStream)
	}
	close(done)
	// 模拟正在进行的工作
	time.Sleep(1*time.Second)

}
