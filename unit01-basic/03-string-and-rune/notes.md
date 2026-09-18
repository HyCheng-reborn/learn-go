# 03 - string、byte 与 rune

> 本目录是按已学知识点重构的复习例子。

## string 不是“字符数组”这么简单

Go 的 `string` 本质上是一段只读字节序列，通常承载 UTF-8 文本。

```go
len(text)
```

返回的是**字节数**，不是“人眼看到的字符数”。

## rune

`rune` 实际是 `int32` 的别名，常用来表示一个 Unicode code point。

```go
[]rune(text)
```

可以把 UTF-8 字符串按 Unicode code point 拆开。

## range 字符串

```go
for i, r := range text
```

其中：

- `i` 是这个 rune 在原字符串中的**字节偏移量**；
- `r` 是解码后的 rune。

这和 Python：

```python
for ch in text:
```

表面相似，但 Go 的 index 是 UTF-8 字节位置，这一点很重要。
