# unit03-struct-method-interface

struct、method、interface、method set、embedding，以及用 interface 做最基础的显式依赖注入。

本单元全部来自 2026-09-18 导入的 ChatGPT 学习归档（见根 README 的「资料来源」），按 ZIP 原目录名归档，仅把 `05-...-IN_PROGRESS` 改成小写 kebab-case。

## 知识点

| 目录 | 状态 | 一句话结论 |
|---|---|---|
| [01-struct-and-constructor](01-struct-and-constructor/) | 📦 归档导入 | Go 没有构造函数，用普通函数 `NewUser` 返回 `*User` 来模拟；值接收者方法 `SayHello` 通过指针也能调 |
| [02-value-vs-pointer-receiver](02-value-vs-pointer-receiver/) | 📦 归档导入 | 值接收者拿到的是副本：`SetAgeByValue(50)` 后 `Age` 仍是 22；指针接收者改原值：`SetAge(99)` 后变 99 |
| [03-interface-method-set](03-interface-method-set/) | 📦 归档导入 | method set 规则：**值接收者 → `T` 和 `*T` 都实现接口；指针接收者 → 只有 `*T` 实现接口，`T` 不实现**。`dog.Speak()` 能调是自动取地址；`var s Speaker = dog` 会编译失败 |
| [04-embedding-composition](04-embedding-composition/) | 📦 归档导入 | embedding 是组合不是 Java 的 `extends`；被嵌入类型的方法会提升，`service.Log(...)`、`car.Start()` 可直接调用 |
| [05-repository-service-di](05-repository-service-di/) | 🚧 进行中 | `UserService` 只依赖 `UserRepository` 接口，不知道具体类型；`NewUserService(repo)` 就是最基础的显式依赖注入。错误用 `fmt.Errorf("get user %d: %w", id, err)` 包装后，`errors.Is` 仍能匹配 `ErrUserNotFound`。加 mock 时踩到的坑：`repo := NewMemoryUserRepository()` 推断出的是**具体类型** `*MemoryUserRepository`，不能再赋 `*MockUserRepository` 给它 |

## 关于 05

归档里这个目录原名 `05-repository-service-di-IN_PROGRESS`，因为「自己从零写 + 运行 + Review + mock」的完整练习流程还没走完。目录名按本仓库命名规则改成小写 kebab-case，**进行中**的状态改记在这里，不写进目录名。代码本身是完整可运行的。

2026-09-18：`MockUserRepository` 由本人独立编写并经 Review 修正（详见该目录 `notes.md` 的「Review 记录」），mock 练习已完成。

尚未做：从零独立实现一遍完整的 repository + service。

## 与 unit02 的关系

`unit02-collections/02-interface-receivers/` 是早期学习路径留下的历史目录，内容和这里的 `02`/`03` 有重叠（同样是 `Dog`/`Speaker` 指针接收者与 method set）。本轮**不移动、不重命名、不删除**它，两者并存：unit02 那份是你亲手写的原始练习，unit03 这份是归档整理版。
