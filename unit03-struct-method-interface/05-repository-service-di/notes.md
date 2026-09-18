# 05 - UserRepository → UserService → DI（进行中）

> 这个单元已经开始讲解，并且你已经能说出核心依赖关系；但还没有完成“自己从零编码 + mock + Review”，所以明确标为 `IN_PROGRESS`。

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

## 当前下一步

这个单元下一步不应该继续让我直接给完整答案，而应该由你自己写一个简单 `FakeUserRepository` 或新的 repository 实现，再让 Qoder 运行和 Review。
