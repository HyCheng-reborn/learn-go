/*
 * @Author: Zhy hycheng368@outlook.com
 * @Date: 2026-09-28 02:44:07
 * @LastEditors: Zhy hycheng368@outlook.com
 * @LastEditTime: 2026-09-28 02:44:29
 * @FilePath: \goproject\unit06-concurrency-and-context\03-buffered-channel\main.go
 * @Description:
 *
 * Copyright (c) 2026 by Zhy, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

func worker(
	id int,
	wg *sync.WaitGroup,
	results chan<- int,
) {
	defer wg.Done()

	res := id * 10 // 算结果
	results <- res // 把结果发送到 results
}

func consume(results <-chan int) {
	for result := range results {
		fmt.Println("result:", result)
	}
}

func main() {
	ch := make(chan int, 3)

	fmt.Println("initial:", len(ch), cap(ch))

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("after send:", len(ch), cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("after receive:", len(ch), cap(ch))

	// 行为 2：close 后还能读已有值
	fmt.Println("\n--- 行为 2：close 后读取 ---")
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	fmt.Println("len after close:", len(ch2))
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2) // channel 已关闭且缓冲数据已读完，因此立即返回 int 零值

	// 行为 3：for range 在 close 后正常结束
	fmt.Println("\n--- 行为 3：for range ---")
	ch3 := make(chan int, 3)
	ch3 <- 10
	ch3 <- 20
	close(ch3)
	for v := range ch3 {
		fmt.Println(v)
	}
	fmt.Println("for range 正常结束")

	//行为4:3worker
	var wg sync.WaitGroup
	results := make(chan int)
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	consume(results)

}
