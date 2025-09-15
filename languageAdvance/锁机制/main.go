package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// calulate()
	calulate2()
}

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

func calulate() {

	res := 0
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lock.Lock()
				res++
				lock.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("结果：", res)
}

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

func calulate2() {

	var res int64 = 0
	var wg sync.WaitGroup
	// var lock sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// lock.Lock()
				atomic.AddInt64(&res, 1)
				// lock.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("结果：", res)
}
