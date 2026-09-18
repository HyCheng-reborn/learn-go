package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u User) SayHello() {
	fmt.Printf("Hello, I am %s, age %d\n", u.Name, u.Age)
}

func main() {
	user := NewUser("Alice", 22)
	user.SayHello()
}
