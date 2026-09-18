package main

import "fmt"

type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("LOG:", msg)
}

type Service struct {
	Logger
}

type Engine struct{}

func (Engine) Start() {
	fmt.Println("engine started")
}

type Car struct {
	Engine
	Name string
}

func main() {
	service := Service{}
	service.Log("hello")

	car := Car{Name: "GoCar"}
	car.Start()
}
