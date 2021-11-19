package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//go sayHello()
	//
	//go func() {
	//	fmt.Println("Hello")
	//}()
	//
	//sayHello:= func() {
	//	fmt.Println("Hello")
	//}
	//go sayHello()


	//var wg sync.WaitGroup
	//sayHello:= func() {
	//	defer wg.Done()
	//	fmt.Println("Hello")
	//}
	//wg.Add(1)
	//go sayHello()
	//wg.Wait()

	// 输出结果welcome
	//var wg sync.WaitGroup
	//salutation:="hello"
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	salutation="welcome"
	//}()
	//wg.Wait()
	//fmt.Println(salutation)

	// 输出三行 good day
	//var wg sync.WaitGroup
	//for _,salutation :=range []string{"hello","greetings","good day"}{
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		fmt.Println(salutation)
	//	}()
	//}
	//wg.Wait()

	// 正常输出 []string内容
	//var wg sync.WaitGroup
	//for _,salutation :=range []string{"hello","greetings","good day"}{
	//	wg.Add(1)
	//	go func(salutation string) {
	//		defer wg.Done()
	//		fmt.Println(salutation)
	//	}(salutation)
	//}
	//wg.Wait()

	// 计算goroutine 大小
	memconsumed:= func() uint64{
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {wg.Done();<-c}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before:=memconsumed()

	for i:=numGoroutines;i>0;i--{
		go noop()
	}

	wg.Wait()
	after:=memconsumed()
	fmt.Printf("%.3fkb",float64(after-before)/numGoroutines/1000)




}

func sayHello(){
	fmt.Println("Hello")
}