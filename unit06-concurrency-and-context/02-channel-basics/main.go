/*
 * @Author: Zhy hycheng368@outlook.com
 * @Date: 2026-09-28 00:05:34
 * @LastEditors: Zhy hycheng368@outlook.com
 * @LastEditTime: 2026-09-28 02:24:33
 * @FilePath: \goproject\unit06-concurrency-and-context\02-channel-basics\main.go
 * @Description:
 *
 * Copyright (c) 2026 by Zhy, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func work(ch chan int) {
	time.Sleep(300 * time.Millisecond)
	ch <- 42
}

func main() {
	ch := make(chan int)
	// ch <- 42

	// fmt.Println(<-ch)

	go work(ch)

	result := <-ch

	fmt.Println("result:", result)
}
