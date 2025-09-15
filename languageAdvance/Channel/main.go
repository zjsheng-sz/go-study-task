package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// channel1()
	channel2()
}

/*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/

func channel1() {

	var sendWg sync.WaitGroup
	var recieveWg sync.WaitGroup

	channel := make(chan int)
	results := []int{}

	recieveWg.Add(1)
	go func() {
		defer recieveWg.Done()
		for res := range channel {
			results = append(results, res)
			fmt.Println("recieve to channel: ", res)
		}
	}()

	for i := 1; i <= 10; i++ {
		sendWg.Add(1)
		go func() {
			defer sendWg.Done()
			channel <- i
			fmt.Println("send to channel: ", i)
			time.Sleep(time.Second * 1)
		}()
	}

	sendWg.Wait()
	close(channel)
	recieveWg.Wait()
}

func channel2() []int {

	channel := make(chan int, 10)
	results := []int{}

	go func() {
		for res := range channel {
			results = append(results, res)
			println("recieve data: ", res)
		}
	}()

	for i := 1; i <= 100; i++ {
		go func() {
			channel <- i
			println("send data: ", i)
			if i == 100 {
				close(channel)
			}
		}()
	}

	timeOut := time.After(1 * time.Second)

	for {
		select {
		case v, ok := <-channel:
			if !ok {
				fmt.Println("Channel已关闭")
				return results
			}
			fmt.Println("接收到:", v)
		case <-timeOut:
			fmt.Println("超时退出")
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)

		}
	}

}
