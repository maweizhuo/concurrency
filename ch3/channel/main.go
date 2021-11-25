package main

import "fmt"

func main()  {
	//stringStream:= make(chan string)
	//go func() {
	//	stringStream<-"Hello channels"
	//}()
	////fmt.Println(<-stringStream)
	//salutaion,ok:=<-stringStream
	//fmt.Printf("(%v)：%v",ok,salutaion)

	//valueStream:=make(chan interface{})
	//close(valueStream)

	//  从已关闭的通道读取。 输出默认值(false): 0
    //intStream:=make(chan int)
    //close(intStream)
    //integer,ok:=<-intStream
    //fmt.Printf("(%v):%v",ok,integer)

     // 1 2 3 4 5
   // intStream:=make(chan int)
   // go func() {
   // 	defer close(intStream)
   // 	for i:=1;i<=5;i++{
   // 		intStream<-i
	//	}
	//}()
   //for integer:=range intStream{
   //	fmt.Printf("%v ",integer)
   //}

   //unblocking goroutines...
	//0 has begun
	//1 has begun
	//2 has begun
	//3 has begun
	//4 has begun
   //begin:=make(chan  interface{})
   //var wg sync.WaitGroup
   //for i:=0;i<5;i++{
   //	wg.Add(1)
   //	go func(i int) {
   //		defer wg.Done()
   //		<-begin // begin通道进行读取，由于通道中没有任何值，会产生阻塞
   //		fmt.Printf("%v has begun \n",i)
	//}(i)
   //}
   //fmt.Println("unblocking goroutines...")
   //close(begin)  // 关闭通道，这样所有goroutine的阻塞会被解除。
   //wg.Wait()

   //var stdoutBuff bytes.Buffer
   //defer stdoutBuff.WriteTo(os.Stdout)
   //intStream:=make(chan int,4)
   //go func() {
   //	defer close(intStream)
   //	defer fmt.Fprintln(&stdoutBuff,"Producer Done.")
   //	for i:=0;i<5;i++{
   //		fmt.Fprintf(&stdoutBuff,"Sending: %d\n",i)
   //		intStream<-i
	//}
   //}()
   //for integer:=range intStream{
   //	fmt.Fprintf(&stdoutBuff,"Received %v .\n",integer)
   //}

   chanOwner:= func()<-chan int {
        resultStream:=make(chan int,5)
        go func() {
        	defer close(resultStream)
        	for i:=0;i<=5;i++{
        		resultStream<-i
			}
		}()
        return resultStream
   }
   resultStream:=chanOwner()
   for result:=range resultStream{
   	fmt.Printf("Received : %d\n",result)
   }
   fmt.Println("Done receiving !")


}

