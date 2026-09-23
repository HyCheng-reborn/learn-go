# Unit 05 · 02 Table-Driven Tests 学习笔记

## 1. 本节目标

本节学习 Go 常见的表驱动测试（table-driven tests），并把多个相似测试场景整理成“测试数据表 + 一套公共测试逻辑”。

重点：

- `[]struct{...}`
- `for _, tt := range tests`
- `t.Run(...)`
- `errors.Is(...)`
- `reflect.DeepEqual(...)`
- statement coverage
- coverage 与行为覆盖的区别

---

## 2. 为什么需要表驱动测试

当多个测试只有输入和预期结果不同，而测试流程基本一致时，可以把变化的数据放进测试表。

例如：

```text
success
not found
database error
```

不需要分别重复写三套：

```text
创建 fake
创建 service
调用 GetUser
检查结果
```

而是写成：

```text
测试数据表
+
一套公共测试逻辑
```

---

## 3. 测试表

典型结构：

```go
tests := []struct {
    name     string
    repoUser *User
    repoErr  error
    wantUser *User
    wantErr  error
}{
    // cases
}
```

`[]struct{...}` 表示“匿名 struct 组成的 slice”。

每个元素就是一条 test case。

---

## 4. 当前三个测试场景

### success

```text
Repository:
user = Alice
err  = nil

期望：
user = Alice
err  = nil
```

### not found

```text
Repository:
user = nil
err  = ErrUserNotFound

期望：
user = nil
errors.Is(err, ErrUserNotFound) == true
```

### database error

```text
Repository:
user = nil
err  = ErrDatabase

期望：
user = nil
errors.Is(err, ErrDatabase) == true
```

---

## 5. `for _, tt := range tests`

```go
for _, tt := range tests {
    ...
}
```

含义：

- `_`：忽略索引
- `tt`：当前 test case 的值

当前元素是 struct，因此每次迭代会把当前 struct 的值复制给 `tt`。

如果 struct 内包含 pointer，例如：

```go
repoUser *User
```

复制的是 pointer value，不会自动深拷贝其指向的对象。

---

## 6. `t.Run` 与 subtest

```go
t.Run(tt.name, func(t *testing.T) {
    ...
})
```

会为每条 case 创建一个命名子测试。

输出类似：

```text
TestServiceGetUser
├── success
├── not_found
└── database_error
```

优点：

- 很容易知道哪条 case 失败
- 可以单独运行某个 subtest
- 多条 case 可以复用同一套测试逻辑

---

## 7. Fake Repository

```go
type fakeRepository struct {
    user *User
    err  error
}

func (r *fakeRepository) FindByID(id int64) (*User, error) {
    return r.user, r.err
}
```

每条 test case 通过：

```text
repoUser
repoErr
```

控制 Fake Repository 的行为。

这使 Service 测试无需真实数据库。

---

## 8. 检查“有没有 error”

当前测试使用：

```go
if (err != nil) != (tt.wantErr != nil) {
    t.Fatalf(...)
}
```

它比较：

```text
实际是否有 error
```

与：

```text
预期是否有 error
```

是否一致。

逻辑：

```text
实际无错 + 预期无错 → 继续
实际有错 + 预期有错 → 继续检查错误根因
实际有错 + 预期无错 → FAIL
实际无错 + 预期有错 → FAIL
```

---

## 9. 检查错误根因

```go
if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
    t.Errorf(...)
}
```

因为 Service 使用 `%w` 包装 error：

```go
fmt.Errorf("get user %d: %w", id, err)
```

所以测试不能只依赖：

```go
err == tt.wantErr
```

而应该用：

```go
errors.Is(err, tt.wantErr)
```

沿 error chain 检查根因。

---

## 10. `reflect.DeepEqual`

当前测试：

```go
reflect.DeepEqual(got, tt.wantUser)
```

可以比较两个 `*User` 指向的数据内容，而不只是比较 pointer 地址。

对于简单业务对象，显式比较字段也很常见：

```go
got.ID
got.Name
```

这样失败时更容易知道具体哪个字段不正确。

---

## 11. Statement Coverage

执行：

```powershell
go test -cover ./unit05-testing/02-table-driven-tests
```

当前结果：

```text
coverage: 100.0% of statements
```

这表示被统计的 statement 都至少执行过一次。

---

## 12. Coverage 不等于测试质量

即使 coverage 是 100%，仍然可能有 bug。

Coverage 只能回答：

> 哪些代码执行过？

不能回答：

> 执行后的业务行为有没有被正确验证？

例如业务代码错误地把：

```go
usr.Name = ""
```

写入成功路径，这行代码本身仍然可以被执行。

如果测试根本不验证 Name，coverage 仍可能是 100%，但 bug 没被发现。

---

## 13. 当前实际覆盖的三种业务行为

### 1. 成功找到用户

```text
Repository 成功
→ Service 返回 User
→ err == nil
```

### 2. 未找到用户

```text
Repository 返回 ErrUserNotFound
→ Service 返回 nil user
→ error chain 保留 ErrUserNotFound
```

### 3. 数据库出错

```text
Repository 返回 ErrDatabase
→ Service 返回 nil user
→ error chain 保留 ErrDatabase
```

相比单纯说：

```text
coverage = 100%
```

描述这些行为更能说明测试实际验证了什么。

---

## 14. 单元测试边界

当前真正的 unit under test 是：

```text
Service.GetUser
```

Fake Repository 只是测试辅助设施。

以后可以分层测试：

```text
Service test
→ Fake Repository
→ 测业务逻辑

Repository test
→ 测 SQL / 数据映射

Handler test
→ Fake Service
→ 测 HTTP 行为
```

不同测试关注不同职责。

---

## 15. Python / Java 对照

Python pytest 常用参数化测试：

```python
@pytest.mark.parametrize(...)
```

Java JUnit 有：

```text
ParameterizedTest
```

Go 常直接使用：

```text
slice
+
anonymous struct
+
for range
+
t.Run
```

这是典型的 Go 风格：用普通语言结构构建测试。

---

## 16. 本节核心结论

1. 表驱动测试把“变化的数据”和“重复的测试逻辑”分离；
2. `[]struct{...}` 可以表示测试用例表；
3. `for range` 遍历各条 case；
4. `t.Run` 为每条 case 建立命名 subtest；
5. `errors.Is` 用于验证包装后的错误根因；
6. statement coverage 只说明代码执行过；
7. 100% coverage 不等于没有 bug；
8. 更重要的是明确哪些业务行为与边界被验证；
9. 当前 `GetUser()` 已覆盖成功、NotFound、Database Error 三种核心行为。
