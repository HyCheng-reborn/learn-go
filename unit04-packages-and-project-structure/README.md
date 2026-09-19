# unit04-packages-and-project-structure

把单文件里的类型、接口、实现、Service 拆分到不同 package，理解 package、import、exported/unexported、module import path、package dependency 和 main wiring。

本 unit 是**本人手写练习**（非 2026-09-18 的 ChatGPT 归档导入）：`main.go` 里先写编号注释要求、再由 Review 填代码。

## 知识点

| 目录 | 状态 | 一句话结论 |
|---|---|---|
| [01-package-import-export](01-package-import-export/) | ✅ | 同一个 module 内按 `learn-go/<unit>/<example>/<pkg>` 的完整 import path 引用子包；`service` 包用到别的包的类型必须显式 import 并加包名前缀（`repository.UserRepository`、`*model.User`），否则编译期 `undefined` |
| [02-package-design-and-internal](02-package-design-and-internal/)（2026-09-19） | ✅ | 按业务组织：model/repository/service 合进一个 `user` package（同包多文件互引不需要包名前缀）；`internal` 是 Go 工具链真正执行的访问限制（越界 import 报 `use of internal package not allowed`，已实验验证），`cmd/` 只是工程约定不是限制；接口按消费者需求设计、保持小（`Repository` 只要 `FindByID`，实现可以多于接口如 `Count`）；命名避免 stutter（`user.Service` 而非 `user.UserService`）；局部变量不能和包名撞（`usr` 而非 `user`，否则遮蔽包名编译失败） |

## package 要点速查

- **package 名 ≠ 目录名可随意**：约定每个目录一个 package，包名取目录语义（`model` / `repository` / `service`），import path 则由「module 名 + 目录路径」拼成，与 package 声明名不一定相同。
- **import 的是路径，使用的是包名**：`import "…/repository"` 之后写 `repository.XXX`，前缀是 package 声明的第一个标识符，不是路径末段的强制结果（这里两者恰好一致）。
- **跨包只认导出标识符**：首字母大写的 `UserRepository`、`ErrUserNotFound`、`NewMemoryUserRepository` 可被 `service` / `main` 引用；未导出的只能在包内用。
- **依赖方向**：`model` ← `repository` ← `service`，`main` 负责组装（wiring）。`repository` 通过 `ErrUserNotFound` 暴露哨兵错误，`service` 用 `%w` 包装后 `main` 仍能 `errors.Is(err, repository.ErrUserNotFound)` 穿透匹配。

## 与 unit03 的关系

`unit03-struct-method-interface/05-repository-service-di` 里 repository / service / User 全挤在一个 `package main`；本 unit 把同一条链路按职责拆成四个包，验证 DI 的接口边界正好也是包边界。
