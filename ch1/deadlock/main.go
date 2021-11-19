package main

import (
	"fmt"
	"sync"
	"time"
)



type value struct {
	mu sync.Mutex
	value int
}

// 死锁是所有并发进程都在彼此等待的状态。 在这种情况下，如果没有外部干预，程序将永远不会恢复。
func main() {
	var wg sync.WaitGroup
	printSum:= func(v1,v2 *value) {
		defer wg.Done()
		v1.mu.Lock()  // 1 这里我们试图访问带锁的部分
		defer v1.mu.Unlock() // 2 这里我们试图调用defer关键字释放锁

		time.Sleep(2*time.Second) // 3 这里我们添加休眠时间 以造成死锁
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n",v1.value+v2.value)
	}

	var a,b value
	wg.Add(2)
	go printSum(&a,&b)
	go printSum(&b,&a)
	wg.Wait()

}
