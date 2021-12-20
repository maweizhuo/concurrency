package main

import (
	"fmt"
	"time"
)

func main() {
	// 心跳
	doWork:= func(done <-chan interface{},pulseInterval time.Duration)(<-chan interface{},<-chan time.Time) {
		heartbeat:=make(chan interface{})
		results:=make(chan time.Time)
		go func() {
			//defer close(heartbeat)
			//defer close(results)

			pulse:=time.Tick(pulseInterval)
			workGen:=time.Tick(2*pulseInterval)

			sendPulse:= func() {
				select {
				case heartbeat<- struct {}{}:
				default:
				}
			}
			sendResult:= func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case results<-r:
						return
					}
				}
			}

            // 一直执行
			//for  {
			//	select {
			//	case <-done:
			//		return
			//	case <-pulse:
			//		sendPulse()
			//	case r:=<-workGen:
			//		sendResult(r)
			//	}
			//}
			for i:=0;i<2;i++{
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r:=<-workGen:
					sendResult(r)
				}
			}

		}()
		return heartbeat,results
	}

	done:=make(chan interface{})
	time.AfterFunc(10*time.Second, func() {close(done)})// 我们设置done通道并在10秒后关闭它。

	const timeout=2*time.Second
	heartbeat,results:=doWork(done,timeout/2)
	for{
		select {
		case _,ok:=<-heartbeat:
			if ok==false{
				return
			}
			fmt.Println("pulse")
		case r,ok:=<-results:
			if ok ==false{
				return
			}
	     fmt.Printf("results %v \n",r.Second())
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}

}
