package main

import (
	"fmt"
	"sync"
)

func main() {

	// 输出2次 Creating new instance.
	//myPool:=&sync.Pool{
	//	New: func() interface{} {
	//		fmt.Println("Creating new instance.")
	//		return struct {}{}
	//	},
	//}
	//myPool.Get()
	//instance:=myPool.Get()
	//myPool.Put(instance)
	//myPool.Get()


	//  输出4 calculators were created.
	var numCalcsCreated int
	calcPool:= &sync.Pool{
		New: func() interface{} {
			numCalcsCreated+=1
			mem:=make([]byte,1024)
			return &mem
		},
	}
	// 将池扩展到4k (-- 扩不扩展输出都是4 --)
	//calcPool.Put(calcPool.New())
	//calcPool.Put(calcPool.New())
	//calcPool.Put(calcPool.New())
	//calcPool.Put(calcPool.New())

	const numWorkers  =  1024 *1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i:=numWorkers;i>0;i--{
		go func() {
			defer wg.Done()

			mem:=calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.",numCalcsCreated)

}
