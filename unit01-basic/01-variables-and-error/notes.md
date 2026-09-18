# 01 - 变量、多返回值与 error

## 这个单元学什么

Go 通常不会像 Java 那样依赖异常来表达普通业务失败，而是把错误作为普通返回值显式返回：

```go
value, err := someFunction()
if err != nil {
    // 处理错误
}
```

## 关键例子

```go
var ErrUserNotFound = errors.New("user not found")

func findUser(name string) (string, error) {
    if name == "Alice" {
        return "Alice", nil
    }
    return "", ErrUserNotFound
}
```

这里返回两个值：

```text
正常：("Alice", nil)
失败：(零值, error)
```

## 哨兵错误

`ErrUserNotFound` 是一个可被重复比较的固定错误值。

推荐：

```go
errors.Is(err, ErrUserNotFound)
```

而不是只比较错误字符串。

## 你之前容易忽略的点

如果代码进入：

```go
if err != nil {
    return
}
```

那么这个 `return` 之后的语句不会再执行。

## Java 对照

Java 常见：

```java
throw new UserNotFoundException();
```

Go 常见：

```go
return nil, ErrUserNotFound
```

Go 强迫调用者在调用位置看见错误处理路径。
