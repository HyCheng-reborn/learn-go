package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Println(d.Name, "says woof")
}

func main() {
	rect := Rectangle{
		Width:  3,
		Height: 4,
	}

	circle := Circle{
		Radius: 2,
	}

	printArea(rect)
	printArea(circle)

	//----------------------------------

	dog := Dog{
		Name: "Buddy",
	}

	// 普通方法调用：可以
	dog.Speak()

	// interface：必须传 *Dog
	var s Speaker = &dog
	s.Speak()

	// 下面这一行取消注释后会编译失败
	// var s2 Speaker = dog
}
