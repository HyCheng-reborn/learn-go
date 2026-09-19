# 06 闭卷练习：从零实现 repository + service + DI

这是 `05-repository-service-di` 的「从零独立实现」考核题。**规则：全程不看 05 的代码**，只允许看本文件的题目描述。写完后按「验收」一节自查。通过后可把 05 和 unit03 的状态翻成 ✅。

## 题目（在 main.go 里按 1~7 实现）

1. 定义 `User`：字段 `ID int64`、`Name string`。
2. 定义哨兵错误 `ErrUserNotFound`（`errors.New`，文案 `user not found`）。
3. 定义接口 `UserRepository`，只有一个方法：`FindByID(id int64) (*User, error)`。
4. 实现 `MemoryUserRepository`：内部用 `map[int64]*User`，提供构造函数，预置一条 `1 → Alice`；查不到的 id 返回 `ErrUserNotFound`。
5. 定义 `UserService`：字段类型是**接口** `UserRepository`（不是具体实现类型）；构造函数 `NewUserService` 收接口注入；方法 `GetUser(id int64) (*User, error)` 调用 `repo.FindByID`，出错时用 `fmt.Errorf` 以 `%w` 包装（格式：`get user %d: 原错误`）再返回。
6. 在 `main` 里跑真实实现的两条路径：
   - `GetUser(1)` 成功 → 打印名字；
   - `GetUser(999)` 失败 → 用 `errors.Is` 判定并打印 NotFound，同时打印包装后的完整错误。注意 `errors.Is` 单独判断为 false 时**不能吞掉其它错误**，要留 `err != nil` 兜底分支。
7. 再写一个 `MockUserRepository`（只认 `knownIDs map[int64]bool` 里的 id，其他一律 NotFound），注入**同一个** `UserService` 类型，跑通 `GetUser(100)` 成功、`GetUser(777)` NotFound 两条路径。

## 期望输出（顺序和文案自定，语义对即可；下面是我给的参考）

```
found: Alice
user 999 was not found
wrapped: get user 999: user not found
found: MockUser
mock: user 777 was not found
```

（Mock 的用户名你想叫什么随意，能区分两条路径即可。）

## 验收

- `go vet ./...` 与 `gofmt -l .` 干净；
- `go run ./unit03-struct-method-interface/06-repository-service-from-scratch` 输出符合期望；
- 自查三个考点：
  1. **第 5 步字段/参数类型必须写接口**。如果写成 `repo *MemoryUserRepository`，第 7 步注入 Mock 时会报 `cannot use ... as ... value`——这正是 05 那次 Review 踩的 `:=` 具体类型坑（`mockRepo := &MockUserRepository{}` 本身没问题，问题出在被赋值的一方是不是接口类型）。
  2. `%w` 包装后 `errors.Is(err, ErrUserNotFound)` 仍能穿透匹配；换成 `%v` 或手写 `err.Error()` 比较就断了。
  3. NotFound 分支之后必须有 `case err != nil` 兜底，否则除 NotFound 外的错误会被静默忽略。

## 常见陷阱（只看这一节不算作弊，但建议跑通后再来对）

- `UserService` 的方法是值接收者还是指针接收者要和构造函数返回的类型一致（返回 `*UserService` 就配指针接收者）。
- map 里存的 `*User` 和返回 `*User` 是同一份数据，Mock 或测试时注意别名。
- 忘记 `errors` / `fmt` import。
