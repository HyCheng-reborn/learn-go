# 04 - Pointer：& 与 *

> 本目录是根据已学 pointer 知识重构的复习例子。

## 取地址

```go
p := &n
```

`p` 保存 `n` 的地址。

## 解引用

```go
*p
```

表示“这个地址指向的那个值”。

因此：

```go
*p = 20
```

会直接修改原变量 `n`。

## 传指针

```go
func addOne(n *int) {
    *n = *n + 1
}
```

调用：

```go
addOne(&n)
```

函数拿到的不是 `n` 的一份整数副本，而是 `n` 的地址，因此可以修改原值。

这个知识会直接进入后面的 pointer receiver：

```go
func (u *User) SetAge(age int)
```
