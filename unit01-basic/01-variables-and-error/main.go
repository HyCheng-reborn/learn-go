package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func main() {
	name := "Alice"
	age := 22
	height := 1.68
	isStudent := true

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("height:", height)
	fmt.Println("student:", isStudent)

	age = 23

	fmt.Println("next age:", age)
	fmt.Printf("max:%d\n", max(8, 5))
	n := 4
	fmt.Printf("%d isEven:%t\n", n, isEven(n))
	n = 5
	fmt.Printf("%d isEven:%t\n", n, isEven(n))
	if result, err := divide(10, 2); err == nil {
		fmt.Printf("divide result: %d\n", int(result))
	} else {
		fmt.Println("cannot divide by zero")
	}
	if result, err := divide(10, 0); err == nil {
		fmt.Printf("divide result: %d\n", int(result))
	} else {
		fmt.Println("cannot divide by zero")
	}
	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Printf("divide result: %d\n", int(result))

	// age, err := getAge(-1)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Printf("age: %d\n", age)

	user, err := findUser("Bob")
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("user not found")
			return
		}
	}
	fmt.Printf("user: %s\n", user)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isEven(n int) bool {
	return n%2 == 0
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func getAge(age int) (int, error) {
	if age < 0 {
		return 0, errors.New("age cannot be negative")
	}
	return age, nil
}

func findUser(name string) (string, error) {
	if name != "Alice" {
		return "", ErrUserNotFound
	}
	return name, nil
}
