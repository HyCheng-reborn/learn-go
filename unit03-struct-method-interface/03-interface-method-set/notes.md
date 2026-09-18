# 03 - Interface、Method Set 与 Dog / *Dog

这是当前阶段最关键、也最容易混淆的知识点之一。

## interface 是什么

```go
type Speaker interface {
    Speak()
}
```

它的意思不是：

> 必须声明 implements Speaker。

而是：

> 只要某个类型的方法集合满足 `Speak()`，它就实现 `Speaker`。

Go 是隐式接口实现。

## 本例

```go
type Dog struct {
    Name string
}

func (d *Dog) Speak() {
    fmt.Println(d.Name, "says woof")
}
```

这里 `Speak` 是 pointer receiver。

因此：

| 方法定义 | `Dog` 实现 Speaker | `*Dog` 实现 Speaker |
|---|---:|---:|
| `func (d Dog) Speak()` | ✅ | ✅ |
| `func (d *Dog) Speak()` | ❌ | ✅ |

本例属于第二行。

## 为什么 `dog.Speak()` 又能调用

```go
dog := Dog{Name: "Buddy"}
dog.Speak()
```

这里是**普通方法调用语法**。

`dog` 是可取地址变量，所以编译器可以自动转换为：

```go
(&dog).Speak()
```

但是：

```go
var s Speaker = dog
```

是在判断：

```text
Dog 这个类型本身是否满足 Speaker？
```

这时按 method set 判断，不会为了让接口赋值通过而把 `Dog` 偷偷改成 `*Dog`。

## 一句话记忆

```text
普通方法调用能否写：有自动取地址/解引用的语法便利。
接口是否实现：严格看类型的 method set。
```

不要把这两个问题混在一起。
