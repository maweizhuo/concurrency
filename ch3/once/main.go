package main

import (
	"fmt"
	"sync"
)

func main() {


	// 输出 count is 1
	//var count int
	//increment:= func() {
	//	count++
	//}
	//var once sync.Once
	//var wg sync.WaitGroup
	//wg.Add(100)
	//for i:=0;i<100;i++{
	//	go func() {
	//		defer wg.Done()
	//		once.Do(increment)
	//	}()
	//}
	//
	//wg.Wait()
	//fmt.Printf("Count is %d \n",count)


	// grep -ir sync.Once $(go env GOROOT)/src |wc -l  // 可以输出go本身使用once次数


	// 输出1
	var count int
	increment:= func() {
		fmt.Println("increment")
		count++}
	decrement:= func() {
		fmt.Println("decrement")
		count--}

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count : %d \n",count)


}
