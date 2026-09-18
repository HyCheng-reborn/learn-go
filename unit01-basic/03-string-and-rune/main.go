package main

import "fmt"

func main() {
	text := "Go语言"

	fmt.Println("text:", text)
	fmt.Println("byte length:", len(text))
	fmt.Println("rune count:", len([]rune(text)))

	fmt.Println("range over string:")
	for byteIndex, r := range text {
		fmt.Printf("byteIndex=%d rune=%c unicode=%U\n", byteIndex, r, r)
	}
}
