# 02 - Value Receiver vs Pointer Receiver

## value receiver

```go
func (u User) SayHello()
```

接收者 `u` 是原值的一份复制。

如果在方法里改字段：

```go
func (u User) SetAgeByValue(age int) {
    u.Age = age
}
```

改的是副本，不会持久修改原 `user`。

## pointer receiver

```go
func (u *User) SetAge(age int) {
    u.Age = age
}
```

接收者是指向原对象的指针，因此可以修改原值。

## 很重要：方法调用的语法糖

如果 `user` 是可取地址变量，下面这句：

```go
user.SetAge(99)
```

即使 `SetAge` 是 `*User` receiver，也可以调用。

编译器可以把它理解为：

```go
(&user).SetAge(99)
```

但这个“自动取地址”规则**不能直接推出** `User` 实现了某个要求该方法的 interface。

接口实现要看 method set，下一单元就是这个区别。
