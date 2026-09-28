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

import "fmt"

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
}
