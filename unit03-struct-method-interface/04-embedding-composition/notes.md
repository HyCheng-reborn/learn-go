# 04 - Embedding / Composition / Method Promotion

## embedding

```go
type Service struct {
    Logger
}
```

这里没有字段名，只有类型名 `Logger`，称为 anonymous field / embedded field。

于是可以直接写：

```go
service.Log("hello")
```

看起来像 `Service` 自己有 `Log` 方法。

## 但它不是 Java extends

更准确的关系是：

```text
Service has a Logger
Car     has an Engine
```

而不是：

```text
Service is a Logger
Car     is an Engine
```

Go 这里强调的是 composition。

## Method Promotion

`Logger` 的方法可以被“提升”到 `Service` 的选择器上：

```go
service.Log(...)
```

实际上底层仍然是 embedded field 提供的方法。

## Java 对照

Java inheritance：

```java
class Car extends Engine
```

表达 is-a。

Go embedding：

```go
type Car struct {
    Engine
}
```

更接近组合，只是提供了方便的方法提升语法。
