# Go Package / Import / Export 学习笔记

## 1. 本节目标

把原本全部放在一个 `main.go` 中的：

- `User`
- `UserRepository`
- `MemoryUserRepository`
- `UserService`
- `main`

拆分到不同 package，并理解：

- package
- import
- exported / unexported
- module import path
- package dependency
- main wiring

---

## 2. 当前目录

```text
01-package-import-export/
├── main.go
├── notes.md
├── model/
│   └── user.go
├── repository/
│   ├── user_repository.go
│   └── memory.go
└── service/
    └── user_service.go
```

---

## 3. Go 中 package 的边界

Go 中应该优先理解为：

> 一个目录通常对应一个 package。

例如：

```text
repository/
├── user_repository.go
└── memory.go
```

两个文件都声明：

```go
package repository
```

因此它们属于同一个 package。

同 package 内的代码可以直接使用彼此的标识符，不需要 import 自己。

例如 `memory.go` 可以直接使用：

```go
ErrUserNotFound
```

---

## 4. 跨 package 使用名称

跨 package 时需要：

```go
package.Identifier
```

例如：

```go
model.User
repository.UserRepository
repository.NewMemoryUserRepository()
service.NewUserService(...)
```

而同 package 内不需要 package 前缀。

---

## 5. Import path 从哪里来

根 `go.mod`：

```go
module learn-go
```

因此：

```text
C:\goproject\
unit04-packages-and-project-structure\
01-package-import-export\
model
```

对应的 import path 是：

```go
learn-go/unit04-packages-and-project-structure/01-package-import-export/model
```

Import 使用的是：

> module path + module 内相对目录

不是 Windows 文件系统绝对路径。

---

## 6. Export 规则

Go 不使用 Java 的：

```text
public
private
protected
```

而通过标识符首字母大小写控制 package 可见性。

例如：

```go
User
NewUserService
ErrUserNotFound
```

首字母大写：

```text
exported
→ 其他 package 可以使用
```

例如：

```go
repository.NewMemoryUserRepository()
```

而：

```go
repo
users
```

首字母小写：

```text
unexported
→ 只能在当前 package 中直接访问
```

例如：

```go
type UserService struct {
	repo repository.UserRepository
}
```

外部 package 不能直接操作 `repo`。

可以粗略类比 Java：

```text
Go 大写开头 ≈ public
Go 小写开头 ≈ package 内部实现细节
```

但二者的访问控制机制并不完全等价。

---

## 7. UserService 的跨包依赖

```go
type UserService struct {
	repo repository.UserRepository
}
```

表示：

> `UserService` 依赖 `repository` package 中定义的 `UserRepository` 接口。

不是依赖具体的：

```go
*repository.MemoryUserRepository
```

构造：

```go
func NewUserService(
	repo repository.UserRepository,
) *UserService {
	return &UserService{repo: repo}
}
```

依然是依赖注入。

---

## 8. 返回其他 package 的类型

`User` 已经属于 `model` package，因此：

```go
func (s *UserService) GetUser(id int64) (*model.User, error)
```

不能再直接写：

```go
*User
```

除非 `User` 就定义在当前 package。

---

## 9. main 的职责：wiring

`main.go`：

```go
repo := repository.NewMemoryUserRepository()
userService := service.NewUserService(repo)
```

作用是：

```text
创建具体 Repository
        ↓
把 Repository 注入 Service
        ↓
调用 Service
```

即：

```text
main
  ↓
repository implementation
  ↓
service
```

这叫 manual wiring。

与 Spring 相比：

```text
Spring：
IoC Container 自动创建和注入

Go 当前写法：
main 中手动创建和注入
```

核心的 DI 思想没有变化。

---

## 10. 当前依赖关系

```text
main
 ├── repository
 └── service
       ├── repository
       └── model

repository
 └── model
```

Go 禁止 import cycle，因此 package 依赖应尽量保持单向。

例如不能：

```text
service
   ↓
repository
   ↓
service
```

否则无法编译。

---

## 11. 本节容易混淆的点

### 文件不等于 package

Java 容易形成：

```text
一个文件 ≈ 一个 class
```

Go 更应该理解：

```text
一个目录 ≈ 一个 package
```

同一 package 可以由多个 `.go` 文件组成。

### interface 仍然有效

代码拆 package 后：

```go
repository.UserRepository
```

仍然只是接口要求。

`*repository.MemoryUserRepository` 因为实现了：

```go
FindByID(id int64) (*model.User, error)
```

所以仍然隐式满足接口。

### 包拆开以后只是名字的位置变化

原来的：

```go
User
UserRepository
NewUserService()
```

变成：

```go
model.User
repository.UserRepository
service.NewUserService()
```

核心业务逻辑没有变化。

---

## 12. 本节验证

执行：

```powershell
gofmt -w .\unit04-packages-and-project-structure\01-package-import-export

go test ./unit04-packages-and-project-structure/01-package-import-export/...

go run ./unit04-packages-and-project-structure/01-package-import-export
```

结果：

```text
found: Alice
user 999 was not found
```

所有 package 编译通过。

---

## 13. Java / Spring 对照

Java / Spring：

```text
User
UserRepository
UserService
Spring IoC 自动 wiring
```

当前 Go：

```text
model.User
repository.UserRepository
service.UserService
main 手动 wiring
```

目前最重要的差异：

```text
Java：
public/private 控制访问

Go：
标识符首字母大小写控制 package export


Java：
implements 显式实现接口

Go：
method set 满足接口即可隐式实现


Spring：
IoC Container 自动注入

Go：
当前通过 NewXxx() + main 手动注入
```

---

## 14. 本节结论

本节真正要掌握的不是“把代码拆成多个文件夹”，而是：

1. 目录通常是 package 边界；
2. 同 package 多文件共同组成一个包；
3. 跨 package 通过 import 和 `package.Identifier` 使用标识符；
4. 大写开头表示 exported，小写开头表示 package 内部实现；
5. import path 来源于 `module path + module 内目录路径`；
6. `main` 可以负责创建具体实现并进行 manual wiring；
7. Go 禁止 import cycle，因此 package 依赖方向需要保持清晰。
