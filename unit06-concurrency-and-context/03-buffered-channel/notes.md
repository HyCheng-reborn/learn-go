# Unit 06 · 03 Buffered / Closed / Directional Channel 学习笔记

## 1. Buffered Channel

```go
ch := make(chan int, 3)
```

`3` 表示 channel 的缓冲容量为 3，可以暂存最多 3 个 `int`。

无缓冲 channel：

```go
make(chan int)
```

需要发送方与接收方直接配合。

有缓冲 channel：

```go
make(chan int, 3)
```

在缓冲未满时，发送方可以暂时领先接收方。

## 2. `len` 与 `cap`

```go
len(ch)
```

表示当前缓冲区中已有多少个值。

```go
cap(ch)
```

表示缓冲区最大容量。

## 3. `close(ch)`

```go
close(ch)
```

表示以后不会再向该 channel 发送新值。

它不会清空已有缓冲数据。

关闭后：

- 已有数据仍可继续接收；
- 再发送会 panic；
- 已关闭且数据耗尽后，接收会立即得到元素类型零值。

若需要区分“真的发送了零值”和“channel 已关闭耗尽”，使用：

```go
value, ok := <-ch
```

其中：

```text
ok == true
→ 收到正常发送的值

ok == false
→ channel 已关闭且数据已耗尽
```

## 4. `for range channel`

```go
for v := range ch {
    fmt.Println(v)
}
```

会持续接收数据。

只有 channel 已关闭并且剩余数据全部读完，循环才会结束。

## 5. 多 Worker + WaitGroup + Channel

典型结构：

```text
worker 1 ─┐
worker 2 ─┼──→ results ───→ consumer
worker 3 ─┘
```

WaitGroup 等待所有 worker 完成；全部完成后关闭 `results`；consumer 使用 `for range` 读到结束。

常见结构：

```go
go func() {
    wg.Wait()
    close(results)
}()

for result := range results {
    // consume
}
```

## 6. 为什么不能任意 Worker 关闭结果 Channel

某个 worker 过早：

```go
close(results)
```

其他 worker 之后再：

```go
results <- value
```

会：

```text
panic: send on closed channel
```

多个 worker 重复 close 也可能：

```text
panic: close of closed channel
```

因此通常由能确定“所有发送者都完成”的一方负责关闭。

## 7. 为什么 main 不能先 Wait 再读取无缓冲 Channel

如果 main：

```text
wg.Wait()
```

而 worker：

```text
results <- value
```

因为结果 channel 无缓冲，worker 等待接收者；main 又在等 worker `Done()`。

形成：

```text
main 等 worker
worker 等 main
```

因此 deadlock。

## 8. Directional Channel

双向：

```go
chan int
```

只发送：

```go
chan<- int
```

只接收：

```go
<-chan int
```

例如：

```go
func worker(results chan<- int) {
    results <- 42
}
```

```go
func consume(results <-chan int) {
    fmt.Println(<-results)
}
```

## 9. 为什么双向 Channel 能传给单向参数

`chan int` 同时具备发送和接收能力。

传给：

```go
chan<- int
```

时，只把可用能力缩小为“发送”。

传给：

```go
<-chan int
```

时，只把可用能力缩小为“接收”。

仍然操作同一个底层 channel，没有创建新的 channel。

能力可以：

```text
完整 → 缩小
```

不能从一个已经只有发送权限的 channel 凭空恢复完整双向权限。

## 10. 核心结论

1. buffered channel 可以暂存有限数量的数据；
2. `close` 表示以后不再发送，不是清空数据；
3. `for range channel` 依赖 channel 最终关闭才能自然结束；
4. WaitGroup、close、range 各自解决不同问题；
5. `chan<- T` 表示只发送；
6. `<-chan T` 表示只接收；
7. 双向 channel 可以缩小为单向视图，但单向权限不能凭空扩大。
