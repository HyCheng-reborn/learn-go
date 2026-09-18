# 05 - UserRepository → UserService → DI（进行中）

> 这个单元已经开始讲解，并且你已经能说出核心依赖关系。「mock + Review」已于 2026-09-18 完成（见文末「Review 记录」），但还没有完成「自己从零编码」，所以仍标为 `IN_PROGRESS`。

## 先看接口

```go
type UserRepository interface {
    FindByID(id int64) (*User, error)
}
```

`UserService` 的需求不是：

> 我必须拿到 `MemoryUserRepository`。

而是：

> 我需要一个满足 `UserRepository` 接口、能够完成 `FindByID` 的 repo。

## 具体实现

```go
type MemoryUserRepository struct {
    users map[int64]*User
}
```

它实现：

```go
func (r *MemoryUserRepository) FindByID(id int64) (*User, error)
```

因此 `*MemoryUserRepository` 满足 `UserRepository`。

## Service 只依赖接口

```go
type UserService struct {
    repo UserRepository
}
```

这意味着：

```text
UserService
    |
    | 只知道接口
    v
UserRepository
    ^
    |
    +---- MemoryUserRepository
    +---- 以后可以是 DBUserRepository
    +---- 以后测试时可以是 FakeUserRepository / MockRepository
```

## 什么是依赖注入

```go
func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}
```

然后：

```go
repo := NewMemoryUserRepository()
service := NewUserService(repo)
```

这里就是最直接的 DI：

```text
先在外部创建依赖 repo
        ↓
把 repo 传给 service
        ↓
service 不在内部偷偷 new 具体实现
```

你之前的理解已经接近核心：

> Service 想完成工作，需要一个能 `FindByID` 的 repo；这个 repo 是 Service 的依赖，把它传进来才能组装出 Service。

## 为什么这样设计

如果写死：

```go
type UserService struct {
    repo *MemoryUserRepository
}
```

Service 就和内存实现绑死。

写成：

```go
repo UserRepository
```

Service 只关心能力，不关心具体实现。

这也是 Go 接口非常典型的实际用途。

## `%w` 与 errors.Is

Service 可以加上业务上下文：

```go
return nil, fmt.Errorf("get user %d: %w", id, err)
```

`%w` 会保留原始错误链，因此外层仍可以：

```go
errors.Is(err, ErrUserNotFound)
```

识别出底层的 `ErrUserNotFound`。

如果只是 `%v`：

```go
fmt.Errorf("...: %v", err)
```

只是把文本拼进去，不会建立可供 `errors.Is` 追踪的错误链。

## Review 记录（2026-09-18）

自己写了 `MockUserRepository` 并在 `main` 里换掉真实 repo，Review 时发现 4 个问题，其中第 1 个是**编译错误**，其余 3 个能编译但语义有缺陷。

### 1. `:=` 推断出的是具体类型，不是接口类型（编译错误）

```text
main.go:83:9: cannot use &MockUserRepository{} (value of type *MockUserRepository)
              as *MemoryUserRepository value in assignment
```

原来的写法：

```go
repo := NewMemoryUserRepository()   // repo 的静态类型被固定为 *MemoryUserRepository
...
repo = &MockUserRepository{}        // 类型不兼容，编译期拒绝
```

`:=` 的规则是：新变量的静态类型 = 右侧表达式的静态类型。`NewMemoryUserRepository()` 的返回类型是 `*MemoryUserRepository`，所以 `repo` 从此只能装这个具体指针类型，和 `UserRepository` 接口没有关系。

容易混淆的地方：`NewUserService(repo)` 之所以能通过，是因为形参声明为 `UserRepository`，把 `*MemoryUserRepository` 传给接口形参时发生了一次**隐式接口转换**（前提是 `*MemoryUserRepository` 的方法集包含 `FindByID`）。但这次转换只发生在「赋值给接口」的那一刻，**不会回头改变 `repo` 自己的静态类型**。

> 结论：接口的可替换性只存在于**接口类型的变量、形参、字段**里，不存在于 `:=` 推断出的具体类型变量里。`UserService.repo` 声明为 `UserRepository` 就是做对了这一点。

两种改法：

```go
// A：显式声明变量类型为接口（改 1 行）
var repo UserRepository = NewMemoryUserRepository()
repo = &MockUserRepository{}   // 合法

// B：不复用变量，给 mock 单独接线（本次采用）
mockRepo := &MockUserRepository{knownIDs: map[int64]bool{100: true}}
mockService := NewUserService(mockRepo)
```

选 B 的理由：两段接线各自独立、不用解释 `var x 接口 = ...` 这种写法。代价是绕开了「同一个接口变量能装不同实现」这个点，所以把该结论记在这份笔记里。

### 2. `fmt.Println(user)` 打印的是指针

`user` 是 `*User`，`fmt` 对结构体指针用 `&{...}` 形式输出，实际打印 `&{100 MockUser}`，和上面 `fmt.Println("found:", user.Name)` 风格不一致。改成打印 `user.Name`（或 `fmt.Printf("%+v\n", *user)`）。

### 3. 只判断 `errors.Is` 会吞掉其他错误

```go
// 改前：不是 NotFound 的错误会静默消失，而且 %w 包出来的消息从没被打印过
_, err = service.GetUser(999)
if errors.Is(err, ErrUserNotFound) {
    fmt.Println("user 999 was not found")
}

// 改后
_, err = service.GetUser(999)
switch {
case errors.Is(err, ErrUserNotFound):
    fmt.Println("user 999 was not found")
    fmt.Println("wrapped:", err)
case err != nil:
    fmt.Println("unexpected error:", err)
}
```

补充语义：`errors.Is(nil, ErrUserNotFound)` 返回 `false`，所以改前的代码在 `err == nil` 时不会 panic，它的问题只是**无法区分「没找到」和「出了别的错」**。

### 4. mock 对任何 id 都成功，走不到错误路径

改前的 `FindByID` 用传进来的 id 现场造一个用户、永远返回 `nil` 错误，导致 `if err != nil` 是死代码，也演示不出「把依赖抽成接口才方便测试失败场景」这个真正的价值。改成只认 `knownIDs`：

```go
type MockUserRepository struct {
    knownIDs map[int64]bool
}

func (r *MockUserRepository) FindByID(id int64) (*User, error) {
    if !r.knownIDs[id] {
        return nil, ErrUserNotFound
    }
    return &User{ID: id, Name: "MockUser"}, nil
}
```

### 顺带印证 method set

`FindByID` 用的是指针接收者，所以只有 `*MockUserRepository` 实现 `UserRepository`，值类型 `MockUserRepository` 不实现——这就是必须写 `&MockUserRepository{}` 带 `&` 的原因，少写一个 `&` 会是第二个编译错误。

### 修正后的实际输出

```text
found: Alice
user 999 was not found
wrapped: get user 999: user not found
found: MockUser
mock: user 777 was not found
mock wrapped: get user 777: user not found
```

`wrapped:` 那两行第一次让 `%w` 的包装结果可见：外层加了 `get user %d` 的业务上下文，`errors.Is` 仍能穿透匹配到 `ErrUserNotFound`。

## 当前下一步

mock 练习已完成。剩下的是**不看这份代码、从零独立实现一遍完整的 repository + service**（含哨兵错误、`%w` 包装、接口字段、一个 fake 实现），再让 Qoder 运行和 Review。
