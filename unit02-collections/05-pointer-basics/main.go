package main

import "fmt"

func addOne(n *int) {
	*n = *n + 1
}

func main() {
	n := 10
	p := &n

	fmt.Println("n:", n)
	fmt.Println("*p:", *p)

	*p = 20
	fmt.Println("after *p = 20, n:", n)

	addOne(&n)
	fmt.Println("after addOne, n:", n)
}
