package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayHello() {
	fmt.Printf("Hello, I am %s\n", u.Name)
}

func (u User) SetAgeByValue(age int) {
	u.Age = age
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func main() {
	user := User{Name: "Alice", Age: 22}

	user.SayHello()

	user.SetAgeByValue(50)
	fmt.Println("after value receiver:", user.Age)

	user.SetAge(99)
	fmt.Println("after pointer receiver:", user.Age)
}
