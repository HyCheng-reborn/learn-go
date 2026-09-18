# 01 - struct 与构造函数式函数

## struct

```go
type User struct {
    Name string
    Age  int
}
```

Go 没有 Java 那种 `class` 关键字。

数据通常放在 struct 里，行为通过 method 绑定到类型上。

## “构造函数”不是语言特殊机制

```go
func NewUser(name string, age int) *User
```

`NewUser` 只是普通函数，只是 Go 社区经常用 `NewXxx` 命名来表达“创建并初始化一个对象”。

它不是 Java constructor：

```java
new User(...)
```

那种由语言定义的特殊语法。
