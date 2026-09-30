# Unit 06 · 05 Context Cancel / Timeout 学习笔记

## 1. `context.Context` 的用途

主要用于沿调用链传播：

- cancellation
- timeout / deadline
- 少量请求级元数据

当前重点是取消与超时。

## 2. `context.Background()`

```go
ctx := context.Background()
```

常作为根 context，本身不会自动取消或超时。

## 3. `context.WithCancel`

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

调用：

```go
cancel()
```

后，`ctx.Done()` 会解除阻塞，`ctx.Err()` 返回：

```text
context canceled
```

## 4. `ctx.Done()`

```go
ctx.Done()
```

返回：

```go
<-chan struct{}
```

即只接收 channel。

因此很适合与 `select` 配合：

```go
select {
case <-ctx.Done():
    // cancel / timeout
case result := <-resultCh:
    // normal result
}
```

## 5. `context.WithTimeout`

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    300*time.Millisecond,
)
defer cancel()
```

超时后 context 自动结束，`ctx.Err()` 返回：

```text
context deadline exceeded
```

## 6. `WithCancel` 与 `WithTimeout`

```text
WithCancel
→ 主要靠手动 cancel()

WithTimeout
→ 可以手动 cancel()
→ 也会在 timeout 到达时自动取消
```

## 7. 为什么 `WithTimeout` 仍然 `defer cancel()`

即使 timeout 最终会自动发生，也通常：

```go
defer cancel()
```

因为如果函数在 deadline 前提前结束，主动 cancel 可以尽早释放 context 关联的 timer、child context 引用等资源，而不必等到超时自然触发。

可以理解：

```text
timeout
→ 保证“最迟什么时候取消”

defer cancel()
→ 保证“提前结束时及时清理”
```

## 8. WaitGroup 在实验中的作用

```go
wg.Add(1)

go func() {
    defer wg.Done()
    worker(ctx)
}()

wg.Wait()
```

确保 main 不会在 worker 处理完取消信号前提前退出。

## 9. Context 参数惯例

通常写：

```go
func DoSomething(ctx context.Context, ...) error
```

要点：

- `context.Context` 通常是第一个参数；
- 不需要 `*context.Context`；
- `Context` 本身是 interface；
- 一般不要把 Context 长期存入 struct；
- 应沿调用链显式传递。

## 10. 后端中的典型传播

```text
HTTP Request
    ↓
r.Context()
    ↓
Service
    ↓
Repository / DB
    ↓
LLM Client
```

请求取消或超时时，下游可以通过同一个 `ctx` 感知并停止工作。

## 11. 核心结论

1. `ctx.Done()` 返回 `<-chan struct{}`；
2. context 取消后，`Done()` 会解除阻塞；
3. `WithCancel` 主要用于手动取消；
4. `WithTimeout` 支持自动超时取消；
5. 即使用 `WithTimeout`，仍应通常 `defer cancel()`；
6. `ctx.Err()` 可区分 `context.Canceled` 与 `context.DeadlineExceeded`；
7. Context 通常作为第一个参数沿调用链传递；
8. 不要把 context 当普通业务数据长期存进 struct。
