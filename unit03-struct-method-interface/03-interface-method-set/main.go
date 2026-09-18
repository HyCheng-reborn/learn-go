package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Println(d.Name, "says woof")
}

func makeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	dog := Dog{Name: "Buddy"}

	// 普通方法调用：dog 可取地址，所以编译器可以自动使用 &dog。
	dog.Speak()

	// 接口赋值：看 method set。
	var s Speaker = &dog
	s.Speak()

	makeItSpeak(&dog)

	// 下面这行如果取消注释，会编译失败：
	// var s2 Speaker = dog
	// 原因：Speak 的接收者是 *Dog，Dog 的 method set 不包含这个方法。
}
