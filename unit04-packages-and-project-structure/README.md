# unit04-packages-and-project-structure

把单文件里的类型、接口、实现、Service 拆分到不同 package，理解 package、import、exported/unexported、module import path、package dependency 和 main wiring。

本 unit 是**本人手写练习**（非 2026-09-18 的 ChatGPT 归档导入）：`main.go` 里先写编号注释要求、再由 Review 填代码。

## 知识点

| 目录 | 状态 | 一句话结论 |
|---|---|---|
| [01-package-import-export](01-package-import-export/) | ✅ | 同一个 module 内按 `learn-go/<unit>/<example>/<pkg>` 的完整 import path 引用子包；`service` 包用到别的包的类型必须显式 import 并加包名前缀（`repository.UserRepository`、`*model.User`），否则编译期 `undefined` |

## package 要点速查

- **package 名 ≠ 目录名可随意**：约定每个目录一个 package，包名取目录语义（`model` / `repository` / `service`），import path 则由「module 名 + 目录路径」拼成，与 package 声明名不一定相同。
- **import 的是路径，使用的是包名**：`import "…/repository"` 之后写 `repository.XXX`，前缀是 package 声明的第一个标识符，不是路径末段的强制结果（这里两者恰好一致）。
- **跨包只认导出标识符**：首字母大写的 `UserRepository`、`ErrUserNotFound`、`NewMemoryUserRepository` 可被 `service` / `main` 引用；未导出的只能在包内用。
- **依赖方向**：`model` ← `repository` ← `service`，`main` 负责组装（wiring）。`repository` 通过 `ErrUserNotFound` 暴露哨兵错误，`service` 用 `%w` 包装后 `main` 仍能 `errors.Is(err, repository.ErrUserNotFound)` 穿透匹配。

## 与 unit03 的关系

`unit03-struct-method-interface/05-repository-service-di` 里 repository / service / User 全挤在一个 `package main`；本 unit 把同一条链路按职责拆成四个包，验证 DI 的接口边界正好也是包边界。
