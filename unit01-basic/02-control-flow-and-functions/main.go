package main

import "fmt"

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	default:
		return "positive"
	}
}

func main() {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	result, ok := divide(10, 2)
	if ok {
		fmt.Println("10 / 2 =", result)
	}

	fmt.Println("-3 is", classify(-3))
}
