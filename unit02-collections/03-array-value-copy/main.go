package main

import "fmt"

func main() {
	a := [3]int{10, 20, 30}
	b := a

	b[0] = 999

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
