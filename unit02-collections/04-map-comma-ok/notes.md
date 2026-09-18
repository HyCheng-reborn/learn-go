# 03 - Map 基础

> 本目录是根据已学 map 知识重构的复习例子。

## 创建

```go
scores := map[string]int{
    "Alice": 90,
}
```

## 读取时的 comma ok

```go
score, ok := scores["Bob"]
```

`ok` 用来区分：

```text
key 不存在
```

和：

```text
key 存在，但 value 恰好是该类型的零值
```

例如 `map[string]int` 中，不存在的 key 直接读取也会得到 `0`。

因此只看 `score == 0` 不足以判断 key 是否存在。

## 删除

```go
delete(scores, "Alice")
```

删除不存在的 key 也不会报错。

## Java 对照

Go：

```go
value, ok := m[key]
```

Java 常见：

```java
map.containsKey(key)
map.get(key)
```
