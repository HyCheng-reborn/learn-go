# Unit 06 · 02 Channel Basics 学习笔记

## 1. 本节目标

理解 Go channel 的最基础用法：

- 创建 channel
- 发送数据
- 接收数据
- 无缓冲 channel 的阻塞特性
- channel 与 goroutine 的配合
- 基础 deadlock 场景

---

## 2. 创建 channel

```go
ch := make(chan int)
```

表示创建一个只能传递 `int` 的 channel。

---

## 3. 发送与接收

发送：

```go
ch <- 42
```

表示把 `42` 发送到 `ch`。

接收：

```go
result := <-ch
```

表示从 `ch` 接收一个值。

---

## 4. 无缓冲 channel

```go
ch := make(chan int)
```

创建的是无缓冲 channel。

没有接收方时：

```go
ch <- 42
```

发送方会阻塞。

没有发送方时：

```go
<-ch
```

接收方会阻塞。

---

## 5. channel 也会形成同步点

```go
func work(ch chan int) {
	ch <- 42
}

func main() {
	ch := make(chan int)
	go work(ch)
	result := <-ch
	fmt.Println(result)
}
```

如果 worker 还没有发送数据，main 会阻塞在 `<-ch`，直到收到值。

---

## 6. Deadlock 实验

```go
func main() {
	ch := make(chan int)

	ch <- 42

	fmt.Println(<-ch)
}
```

会 deadlock。

原因：

```text
main 执行 ch <- 42
↓
无缓冲 channel 没有接收者
↓
main 被阻塞
↓
后面的 <-ch 永远无法执行
↓
deadlock
```

实际观察：

```text
fatal error: all goroutines are asleep - deadlock!
```

---

## 7. 本节核心结论

1. `ch <- value` 是发送；
2. `<-ch` 是接收；
3. 无缓冲 channel 的发送和接收都可能阻塞；
4. 无缓冲 channel 需要发送方与接收方配合；
5. channel 不仅传递数据，也会形成同步点；
6. 同一个 goroutine 在无缓冲 channel 上先发送、后接收，会因为发送先阻塞而无法走到后面的接收代码。
