# Unit 04 · 02 Package Design 与 internal 学习笔记

## 1. 本节目标

这一节不再只学习“怎么 import”，而是开始理解 Go 项目为什么要这样组织 package。

重点包括：

- `internal/` 的真实作用
- `cmd/` 的工程意义
- 按技术层组织 vs 按业务能力组织
- 同 package 多文件如何协作
- package 命名与避免重复命名（stutter）
- interface 为什么更应该从“消费者需求”出发设计
- 小 interface 的价值
- `method set`、interface、DI、package 如何串起来

---

## 2. 当前目录

```text
02-package-design-and-internal/
├── cmd/
│   └── app/
│       └── main.go
└── internal/
    └── user/
        ├── model.go
        ├── repository.go
        └── service.go
```

其中：

```go
internal/user/model.go
internal/user/repository.go
internal/user/service.go
```

都声明：

```go
package user
```

因此它们共同组成一个 `user` package。

---

## 3. 按技术层组织 vs 按业务组织

教学阶段曾使用：

```text
model/
repository/
service/
```

这种方式是按技术职责分层。

大型项目中也常见另一种组织方式：

```text
internal/
├── user/
├── order/
├── auth/
└── resume/
```

每个业务模块内部再放自己的 model、repository、service、handler。

这种方式更强调：

> 强相关的业务代码尽量放在一起。

---

## 4. `internal/` 不是普通目录名

`internal` 在 Go 工具链中有特殊含义。

例如：

```text
02-package-design-and-internal/
├── cmd/app/
└── internal/user/
```

`cmd/app` 位于 `02-package-design-and-internal` 这棵目录树内，因此可以 import：

```go
learn-go/unit04-packages-and-project-structure/02-package-design-and-internal/internal/user
```

但在其允许范围之外的代码不能 import 该 internal package。

---

## 5. internal 非法 import 实验

在 `C:\goproject` 根目录临时创建：

```go
package main

import (
    "fmt"

    "learn-go/unit04-packages-and-project-structure/02-package-design-and-internal/internal/user"
)

func main() {
    repo := user.NewMemoryRepository()
    fmt.Println(repo)
}
```

执行：

```powershell
go run .\internal_probe.go
```

得到：

```text
use of internal package
learn-go/unit04-packages-and-project-structure/02-package-design-and-internal/internal/user
not allowed
```

说明：

> `internal` 不是“社区建议不要使用”，而是 Go 工具链真正执行的 package 访问限制。

实验后删除临时文件，再执行：

```powershell
go test ./unit04-packages-and-project-structure/02-package-design-and-internal/...
```

重新全部通过。

---

## 6. `cmd/` 与 `internal/` 的区别

### `internal/`

具有 Go 工具链层面的特殊 import 规则。

主要表达：

> 这是项目内部实现，不希望被允许范围之外的代码直接依赖。

### `cmd/`

主要是 Go 项目中的常见工程约定，不是和 `internal` 一样的特殊 import 规则。

典型结构：

```text
cmd/
├── server/
│   └── main.go
├── worker/
│   └── main.go
└── migrate/
    └── main.go
```

代表一个仓库可以构建多个可执行程序。

---

## 7. `main.go` 应尽量保持“薄”

真实项目中的 `main` 通常主要负责：

```text
创建依赖
↓
组装对象
↓
启动程序
```

例如：

```go
repo := user.NewMemoryRepository()
svc := user.NewService(repo)
```

这类对象组装通常称为：

```text
wiring
```

---

## 8. 同 package 多文件的意义

当前：

```text
internal/user/
├── model.go
├── repository.go
└── service.go
```

三个文件虽然物理上分开，但逻辑上共同组成：

```go
package user
```

因此：

`repository.go` 可以直接使用：

```go
*User
```

而不需要：

```go
model.User
```

`service.go` 可以直接使用：

```go
Repository
```

而不需要：

```go
repository.Repository
```

文件拆分主要用于帮助人类组织源码，package 才是 Go 的主要代码边界。

---

## 9. 避免命名重复（stutter）

上一节使用：

```go
repository.UserRepository
service.UserService
```

当代码变成：

```go
package user
```

之后，可以写成：

```go
user.Repository
user.Service
user.MemoryRepository
```

而不一定需要：

```go
user.UserRepository
user.UserService
user.MemoryUserRepository
```

因为 package 名 `user` 已经提供上下文。

Go 中常见：

```text
http.Client
json.Decoder
bytes.Buffer
```

而不是重复成：

```text
http.HTTPClient
json.JSONDecoder
bytes.BytesBuffer
```

这种避免重复命名的思想常称为：

```text
avoid stutter
```

---

## 10. 当前 user package

### User

```go
type User struct {
    ID   int64
    Name string
}
```

### Repository

```go
type Repository interface {
    FindByID(id int64) (*User, error)
}
```

### MemoryRepository

```go
type MemoryRepository struct {
    users map[int64]*User
}
```

### Service

```go
type Service struct {
    repo Repository
}
```

### 构造关系

```text
MemoryRepository
       │
       │ 满足 Repository
       ▼
NewService(repo)
       │
       ▼
Service
```

---

## 11. interface 更应该从“消费者需求”设计

最重要的理解：

```go
type Service struct {
    repo Repository
}
```

真正需要 `Repository` 的是 `Service`。

因此接口表达的是：

> Service 为了完成工作，需要依赖具备什么能力。

例如 Service 只使用：

```go
FindByID(...)
```

那么接口可以保持：

```go
type Repository interface {
    FindByID(id int64) (*User, error)
}
```

即使具体的 `MemoryRepository` 还有：

```text
Save
Delete
Count
Clear
```

也没有必要全部加入接口。

---

## 12. “消费者定义接口”的理解

可以类比甲方和乙方：

> 甲方只定义“我需要你达到什么结果”，并不关心乙方内部如何完成。

对应 Go：

```text
Service
= 需求方 / 消费者

Repository interface
= 能力要求

MemoryRepository / PostgresRepository / MockRepository
= 不同实现方
```

Service 只关心：

```text
FindByID 能不能完成
```

不关心：

```text
数据来自 map
PostgreSQL
Redis
网络
还是测试 fake
```

---

## 13. 为什么偏好小 interface

假设具体实现有：

```text
FindByID
Save
Delete
Count
Clear
```

但 Service 只需要：

```text
FindByID
```

那么最好接口也只要求：

```go
type Repository interface {
    FindByID(id int64) (*User, error)
}
```

好处：

- Service 不依赖无关能力
- Fake / Mock 更容易实现
- 更容易替换实现
- 降低耦合
- 符合接口隔离思想

---

## 14. 实现类型可以比 interface 拥有更多方法

例如：

```go
func (r *MemoryRepository) FindByID(...) ...
func (r *MemoryRepository) Save(...) ...
func (r *MemoryRepository) Delete(...) ...
func (r *MemoryRepository) Count() int ...
```

而接口只有：

```go
type Repository interface {
    FindByID(id int64) (*User, error)
}
```

依然满足接口。

因为 interface 表达的是：

> 至少具备这些方法。

而不是：

> 只能有这些方法。

---

## 15. 不应因为实现新增方法就自动修改 interface

例如 `MemoryRepository` 新增：

```go
func (r *MemoryRepository) Count() int
```

如果 `Service` 根本不用 `Count()`，就不应该为了“让接口和实现完全一样”而修改：

```go
type Repository interface {
    FindByID(...)
}
```

是否修改接口应该取决于：

> 消费者是否真的新增了能力需求。

---

## 16. 与 method set 的关系

例如：

```go
func (r *MemoryRepository) FindByID(...)
```

使用 pointer receiver。

因此满足：

```go
Repository
```

的是：

```go
*MemoryRepository
```

而不是：

```go
MemoryRepository
```

所以：

```go
repo := user.NewMemoryRepository()
```

返回 `*MemoryRepository`，可以传给：

```go
user.NewService(repo)
```

这把之前学过的：

```text
pointer receiver
method set
interface
DI
package
```

串到了一起。

---

## 17. Java / Spring 对照

Java 中常见：

```java
interface UserRepository {
    User findById(long id);
}
```

然后：

```java
class PostgresUserRepository implements UserRepository
```

Go 不需要显式：

```text
implements
```

而是：

```text
具体类型的方法集满足 interface
→ 自动满足接口
```

Go 因此更容易让“消费者后来定义一个小接口”。

第三方实现不需要提前知道该接口存在。

---

## 18. Python 对照

Python：

```python
def get_user(repo, id):
    return repo.find_by_id(id)
```

实际上只要求：

```text
repo 有 find_by_id()
```

这是 duck typing。

Go：

```go
type Repository interface {
    FindByID(id int64) (*User, error)
}
```

也是关注“是否具有所需行为”，但区别是：

```text
Python
→ 主要运行时检查

Go
→ interface + 静态类型
→ 编译时检查
```

---

## 19. 本节最终理解

### Package 组织

不需要机械照搬 Java：

```text
entity/
dao/
service/
controller/
```

Go 项目可以根据业务边界组织：

```text
internal/
├── user/
├── order/
├── auth/
└── resume/
```

### Internal

```text
internal/
```

用于建立项目内部 package 的真实访问边界。

### Interface

interface 应优先描述：

> 消费者真正需要的能力。

而不是：

> 具体实现拥有的全部方法。

### DI

Service 依赖：

```go
Repository
```

具体实现由外部注入：

```go
repo := user.NewMemoryRepository()
svc := user.NewService(repo)
```

### 核心链路

```text
cmd/app
   ↓
internal/user
   ↓
Service
   ↓
Repository interface
   ↓
MemoryRepository
```

---

## 20. 本节验证状态

已验证：

```text
gofmt
✅

go test ./unit04-packages-and-project-structure/02-package-design-and-internal/...
✅

go run ./unit04-packages-and-project-structure/02-package-design-and-internal/cmd/app
✅
```

运行结果：

```text
found: Alice
user 999 was not found
```

同时已通过非法 import 实验确认：

```text
internal package 访问限制真实生效
```
