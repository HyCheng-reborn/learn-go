# Unit 06 · 04 Select Basics 学习笔记

## 1. 本节目标

理解 Go `select` 的核心用途：

- 同时等待多个 channel 操作
- 谁 ready 就处理谁
- 没有 ready case 时阻塞
- `default` 的非阻塞语义
- 使用 `time.After` 实现简单 timeout
- 为后续 `context.Context` 铺路

---

## 2. `select` 与 `switch`

`switch` 主要根据值或条件选择分支：

```go
switch {
case x > 0:
    ...
case x == 0:
    ...
}
```

`select` 主要根据 channel 操作是否可以执行来选择分支：

```go
select {
case v := <-ch1:
    ...
case v := <-ch2:
    ...
}
```

核心区别：

```text
switch
→ 看值 / 条件

select
→ 看 channel 操作是否 ready
```

---

## 3. `select` 的阻塞行为

如果：

```go
select {
case <-ch1:
case <-ch2:
}
```

而两个 channel 都没有可执行的操作，那么当前 goroutine 会阻塞。

直到至少一个 case 变为 ready。

---

## 4. 多个 case 同时 ready

如果多个 case 在 `select` 做选择时都可以执行，Go 不保证按源码顺序选择第一个。

应该理解为：

> 从当前可执行的 case 中选择一个执行。

因此不能依赖：

```text
case 写在前面
→ 一定优先
```

---

## 5. 一次 `select` 只执行一个 case

例如：

```go
select {
case v := <-ch1:
    ...
case v := <-ch2:
    ...
}
```

一次只会选中一个 case。

如果想连续处理多个结果，可以：

```go
for i := 0; i < 2; i++ {
    select {
    case v := <-ch1:
        ...
    case v := <-ch2:
        ...
    }
}
```

---

## 6. 两个 worker 的实验

```go
go sendAfter("worker A", 200*time.Millisecond, ch1)
go sendAfter("worker B", 500*time.Millisecond, ch2)
```

使用：

```go
for i := 0; i < 2; i++ {
    select {
    case msg := <-ch1:
        fmt.Println("received:", msg)
    case msg := <-ch2:
        fmt.Println("received:", msg)
    }
}
```

通常会先收到较早 ready 的 worker A，再收到 worker B。

重点不是固定顺序，而是：

> main 不需要预先决定“先等 ch1 还是先等 ch2”。

---

## 7. `default`

```go
select {
case v := <-ch:
    fmt.Println(v)
default:
    fmt.Println("nothing ready")
}
```

如果没有任何 channel case 可以立即执行，则马上走 `default`。

所以：

```text
无 default
→ 可能阻塞

有 default
→ 可以进行非阻塞检查
```

不要在无限循环中滥用空 `default`，否则可能形成 busy loop，持续消耗 CPU。

---

## 8. `time.After` 与 timeout

```go
select {
case result := <-resultCh:
    fmt.Println(result)

case <-time.After(300 * time.Millisecond):
    fmt.Println("timeout")
}
```

`time.After(...)` 会返回一个接收型 channel。

当指定时间到达后，该 channel 会变为 ready。

因此这里相当于：

```text
resultCh 先 ready
→ 正常处理结果

300ms timeout 先 ready
→ timeout
```

---

## 9. 100ms / 1s worker

如果 worker：

```text
100ms 后返回
```

而 timeout：

```text
300ms
```

通常会先走结果 case。

如果 worker：

```text
1s 后返回
```

而 timeout：

```text
300ms
```

会先走 timeout case。

---

## 10. 一个重要隐患：timeout 不等于取消 worker

简单写：

```go
select {
case result := <-resultCh:
    ...
case <-time.After(300 * time.Millisecond):
    ...
}
```

只表示：

> 当前等待者不再等结果了。

它并不会自动通知后台 goroutine：

> 你应该停止工作。

如果后台 goroutine 之后仍然：

```go
resultCh <- result
```

而已经没有接收者，使用无缓冲 channel 时它可能永久阻塞。

这正是下一节 `context.Context` 要解决的问题之一：

```text
超时 / 请求取消
↓
把取消信号传播给后台工作
↓
让 goroutine 主动停止
```

---

## 11. 本节核心结论

1. `select` 用来等待多个 channel 操作；
2. 没有 ready case 且无 `default` 时，当前 goroutine 阻塞；
3. 多个 case 同时 ready 时，不应依赖源码顺序；
4. 一次 `select` 只执行一个 case；
5. `default` 可以实现非阻塞检查，但可能导致 busy loop；
6. `time.After` 可以参与 `select` 实现简单 timeout；
7. timeout 只结束“等待”，不会自动取消后台 goroutine；
8. 真正的取消与超时传播需要 `context.Context`。
