package main

import (
	"fmt"
	"time"
)

func main() {
	// 与switch不同的是，case分支不会被顺序测试，如果没有任何分支的条件可供满足，select会一直等待直到某个case语句完成。
	//var c1, c2 <-chan interface{}
	//var c3 chan<- interface{}
	//select {
	//case <-c1:
	//// do something
	//case <-c2:
	//// do something
	//case c3<- struct {}{}:
	//	// do something
	//}

   //start:=time.Now()
   //c:=make(chan interface{})
   //go func() {
   //	time.Sleep(5*time.Second)
   //	close(c)
   //}()
   //fmt.Println("Blocking on read ...")
	//select {
	//case <-c:
	//	fmt.Printf("unblocked %v later.\n",time.Since(start))
   //}

   //
  //c1:=make(chan interface{})
  //close(c1)
  //c2:=make(chan interface{})
  //close(c2)
  //var c1Count,c2Count int
  //for i:=1000;i>=0;i--{
	//  select {
  //	   case <-c1:
  //	   	c1Count++
	//  case <-c2:
	//	  c2Count++
	// }
  //}
  //fmt.Printf("c1Count: %d\nc2Count: %d\n",c1Count,c2Count)

  // 超时机制
  //var c <-chan int
	//select {
	//case <-c:
	//case <-time.After(1*time.Second):
	//	fmt.Println("Timed out .")
  //}

  // default 无操作的话默认执行语句  In default after 4.208µs
  //start:=time.Now()
  //var c1,c2 chan int
	//select {
	//case <-c1:
	//case <-c2:
	//default:
	//	fmt.Printf("In default after %v \n\n",time.Since(start))
  //}

  // Achieved 5 cycles of work before signalled to stop.
   done:=make(chan interface{})
   go func() {
   	time.Sleep(5*time.Second)
   	close(done)
   }()
	workCounter := 0
 loop:
	for {
		select {
		case <-done:
         break loop
		default:

		}

		workCounter++
		time.Sleep(1*time.Second)
	}
 fmt.Printf("Achieved %v cycles of work before signalled to stop.\n",workCounter)

}
