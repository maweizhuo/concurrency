package main

import (
	"fmt"
	"sync"
	"time"
)
// 饥饿是指并发进程无法获得执行工作所需的任何资源的情况。，
// 饥饿通常意味着有一个或多个贪婪的并发进程不公平地阻止一个或多个并发进程尽可能有效地完成工作，或者根本不可能完成工作。
func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime =1*time.Second
	greedyWorker:= func() {
		defer wg.Done()
		var count int
		for begin :=time.Now();time.Since(begin)<=runtime;{
			sharedLock.Lock()
			time.Sleep(3*time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops \n",count)
	}

    politeWorker:= func() {
    	defer wg.Done()
    	var count int
    	for begin:=time.Now();time.Since(begin)<=runtime;{
    		sharedLock.Lock()
    		time.Sleep(1*time.Nanosecond)
    		sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

    		count++

		}
    	fmt.Printf("Polite worker was able to execute %v work loops \n",count)
	}

    wg.Add(2)
    go greedyWorker()
    go politeWorker()
    wg.Wait()

}
