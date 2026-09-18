# 当前 Go 学习进度

## 已掌握或基本掌握

- Go 变量与基本函数写法
- 多返回值
- Go 的显式 `error` 返回方式
- 哨兵错误 `errors.New(...)`
- `errors.Is`
- array 的值复制
- slice 的 `len` / `cap` / `append`
- slice 共享底层数组以及扩容后可能脱离旧数组
- map 基础操作
- pointer 基础
- `struct`
- method
- value receiver / pointer receiver
- interface 是“满足方法集合即可”，不要求显式 implements
- method set：pointer receiver 时 `T` 与 `*T` 的接口实现差异
- 普通方法调用的自动取地址，与“是否实现接口”是两套规则
- embedding 是组合，不是 Java `extends`

## 当前正在学

### interface + DI

正在理解这条关系：

```text
UserService
    |
    | 依赖一个“能 FindByID 的东西”
    v
UserRepository interface
    ^
    |
MemoryUserRepository
```

核心理解已经形成：

- `UserService` 不需要知道 repository 的具体类型；
- 它只要求 `repo` 满足 `UserRepository` 接口；
- `repo` 是构造 `UserService` 所需的依赖；
- 把具体 repo 传给 `NewUserService`，就是最基础、最显式的一种依赖注入。

但这个单元尚未完成“自己从零写 + 运行 + Review + mock”的完整练习流程，因此归档中标为 `IN_PROGRESS`。
