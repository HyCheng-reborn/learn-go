package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   80,
	}

	scores["Carol"] = 95

	score, ok := scores["Bob"]
	fmt.Println("Bob:", score, "exists:", ok)

	missingScore, ok := scores["Nobody"]
	fmt.Println("Nobody:", missingScore, "exists:", ok)

	delete(scores, "Alice")
	fmt.Println("scores:", scores)
}
