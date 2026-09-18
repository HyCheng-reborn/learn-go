package main

import "fmt"

func main() {
	a := make([]int, 3)

	a[0] = 10
	a[1] = 20
	a[2] = 30

	b := a

	fmt.Println("before")
	fmt.Println("a:", a, "len:", len(a), "cap:", cap(a))
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))

	a = append(a, 40)

	a[0] = 999

	fmt.Println("after")
	fmt.Println("a:", a, "len:", len(a), "cap:", cap(a))
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))
}
