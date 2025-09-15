package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// asyncExecute()

	scheduler := newScheduler(3)
	scheduler.add(func() {
		fmt.Print("task1 start \n")
		time.Sleep(time.Second * 1)
		fmt.Print("task1 end \n")
	})

	scheduler.add(func() {
		fmt.Print("task2 start \n")
		time.Sleep(time.Second * 1)
		fmt.Print("task2 end \n")
	})

	scheduler.add(func() {
		fmt.Print("task3 start \n")
		time.Sleep(time.Second * 1)
		fmt.Print("task3 end \n")
	})

	results := scheduler.run()

	for _, result := range results {
		fmt.Printf("result = %v\n", result)
	}
}

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

func asyncExecute() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
				time.Sleep(time.Second * 1)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
				time.Sleep(time.Second * 1)
			}
		}
	}()

	wg.Wait()
}

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

type Task func()

type TaskResult struct {
	ID       int
	Duration time.Duration
	Error    error
}

type Scheduler struct {
	tasks    []Task
	currency int //最大并发数
}

func newScheduler(concurrent int) *Scheduler {
	return &Scheduler{
		currency: concurrent,
		tasks:    make([]Task, concurrent),
	}
}

func (s *Scheduler) add(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) run() []TaskResult {
	var wg sync.WaitGroup

	taskResult := make([]TaskResult, len(s.tasks))
	channel := make(chan TaskResult, len(s.tasks))
	throttle := make(chan struct{}, s.currency)

	var resultsWg sync.WaitGroup
	resultsWg.Add(1)
	//接收通道值
	go func() {
		defer resultsWg.Done()
		for res := range channel {
			taskResult = append(taskResult, res)
			fmt.Printf("channel 接收到数据 %v\n", res)
		}
	}()

	//执行任务
	for _, task := range s.tasks {

		wg.Add(1)
		throttle <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() {
				<-throttle
			}()

			start := time.Now()
			var err error
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("task panic: %v", r)
				}
			}()

			task()

			channel <- TaskResult{
				ID:       1,
				Error:    err,
				Duration: time.Since(start),
			}
		}()
	}

	wg.Wait()
	close(channel)
    resultsWg.Wait() // 等待结果收集goroutine完成
	close(throttle)
	fmt.Println("run return ", len(taskResult))
	return taskResult
}
