# Unit 06 · 01 Goroutine Basics 学习笔记

## 1. 本节目标

学习 Go 最基础的并发模型：

- goroutine
- `go` 关键字
- `sync.WaitGroup`
- `Add / Done / Wait`
- `defer`
- 多 goroutine 的执行顺序不确定性

---

## 2. goroutine

普通调用：

```go
work()
```

会同步等待 `work()` 返回。

goroutine：

```go
go work()
```

会启动一个新的 goroutine 执行 `work()`，当前 goroutine 不会等待它完成。

---

## 3. main 返回会结束整个程序

仅写：

```go
go work()
```

不能保证 worker 一定执行完成。

如果 main goroutine 很快返回：

```text
main 返回
↓
整个进程结束
```

其他 goroutine 可能还没开始、只执行了一部分，或者刚好执行完。

---

## 4. WaitGroup

`sync.WaitGroup` 可以理解为“未完成任务计数器”。

### Add

```go
wg.Add(1)
```

计数器 +1，表示新增一个要等待的任务。

### Done

```go
wg.Done()
```

计数器 -1，表示一个任务完成。

### Wait

```go
wg.Wait()
```

阻塞当前 goroutine，直到计数器变成 0。

注意：

> `Wait()` 阻塞的是调用它的 goroutine，不一定只能是 main。

---

## 5. 常见正确模式

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()

    // work
}()

wg.Wait()
```

重要顺序：

```text
先 Add
再启动 goroutine
```

不要把 `Add(1)` 放进新 goroutine 后才执行，否则 `Wait()` 可能在任务登记前就看到计数器为 0。

---

## 6. defer

```go
defer wg.Done()
```

表示当前函数返回前执行 `wg.Done()`。

优点是，即使函数中途正常 `return`，`Done()` 仍然会执行。

以后 `defer` 还经常用于关闭文件、释放锁、关闭 HTTP response body 和清理资源。

---

## 7. 多个 goroutine

例如：

```go
wg.Add(3)

go work(0, &wg)
go work(1, &wg)
go work(2, &wg)

wg.Wait()
```

实际输出顺序可能变化：

```text
worker 2: start
worker 0: start
worker 1: start
...
```

创建 goroutine 的代码顺序不等于实际运行顺序。

---

## 8. WaitGroup 只保证“全部完成”，不保证执行顺序

`wg.Add(3)` + 三次 `Done()` 能保证：

```text
所有 worker 完成
↓
Wait() 返回
↓
main 继续
```

但不能保证：

```text
worker 0
↓
worker 1
↓
worker 2
```

依次执行。

因此：

```text
等待全部完成
≠
控制执行顺序
```

---

## 9. Add / Done 不匹配

如果：

```go
wg.Add(3)
```

但只有两个任务最终调用：

```go
wg.Done()
```

那么计数器会停在 1。

`wg.Wait()` 永远等不到 0。

如果此时没有任何 goroutine 能继续推进，Go runtime 通常会检测到 deadlock，并报：

```text
fatal error: all goroutines are asleep - deadlock!
```

---

## 10. 当前实际观察

多次运行后，worker 的启动顺序发生变化。

这说明 goroutine 的调度顺序不能依赖源码启动顺序。

但由于 main 使用 `wg.Wait()`，`main: done` 始终会等所有 worker 完成后才执行。

---

## 11. 本节核心结论

1. `go f()` 启动 goroutine，但不会等待它完成；
2. main 返回会结束整个程序；
3. `WaitGroup` 用于等待一组任务完成；
4. `Add` 增加任务数，`Done` 减少任务数，`Wait` 等到计数为 0；
5. `defer wg.Done()` 是常见安全写法；
6. goroutine 启动顺序不等于运行顺序；
7. `WaitGroup` 只负责“等完”，不负责控制先后顺序；
8. `Add` 和 `Done` 不匹配会导致永久等待，最终可能触发 deadlock。
