package main

import (
	"fmt"
	"time"
)

// sendAfter: 等待 delay 后，把 name 发送到 ch
func sendAfter(name string, delay time.Duration, ch chan<- string) {
	time.Sleep(delay)
	ch <- name
}

func main() {
	// 实验 A：两个 worker，谁先完成先处理
	fmt.Println("=== 实验 A：select 多路复用 ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendAfter("worker A", 200*time.Millisecond, ch1)
	go sendAfter("worker B", 500*time.Millisecond, ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("received:", msg)
		case msg := <-ch2:
			fmt.Println("received:", msg)
		}
	}

	// 实验 B：timeout
	fmt.Println("\n=== 实验 B：timeout ===")
	resultCh := make(chan string)

	// worker：100ms 后发送结果
	go func() {
		time.Sleep(100 * time.Millisecond)
		resultCh <- "result from worker"
	}()

	select {
	case result := <-resultCh:
		fmt.Println(result)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("timeout")
	}
}
