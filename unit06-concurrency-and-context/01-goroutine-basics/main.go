package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d: start\n", id)

	time.Sleep(300 * time.Millisecond)

	fmt.Printf("worker %d: done\n", id)
}

func main() {
	fmt.Println("main: start")

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		i := i
		go work(i, &wg)
	}

	wg.Wait()

	fmt.Println("main: done")
}
