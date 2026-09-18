# 02 - 控制流与函数

> 本目录是根据已经学过的知识点重构的复习例子；历史原始练习文件后来被覆盖，不能保证逐字一致。

## if

Go 的条件不需要括号：

```go
if x > 0 {
    ...
}
```

## for

Go 没有单独的 `while` 关键字，循环统一用 `for`。

```go
for i := 0; i < 10; i++ {
}
```

也可以写成类似 while：

```go
for condition {
}
```

## switch

Go 的 `switch` 默认不会像 C/Java 那样继续向下贯穿，不需要每个 case 都写 `break`。

## 多返回值

```go
func divide(a, b int) (int, bool)
```

这是 Go 很常见的设计：同时返回结果和状态。

后面的 `value, err := ...` 也是同一种语言能力。
