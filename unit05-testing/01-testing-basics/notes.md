# Unit 05 · 01 Go Testing Basics 学习笔记

## 1. 本节目标

本节开始使用 Go 标准库 `testing` 学习单元测试，并把之前学过的 interface、dependency injection、fake repository、error wrapping、`errors.Is` 应用到测试中。

## 2. Go 测试文件和测试函数

Go 测试文件使用：

```text
xxx_test.go
```

测试函数格式：

```go
func TestSomething(t *testing.T) {
}
```

运行：

```powershell
go test ./unit05-testing/01-testing-basics
go test -v ./unit05-testing/01-testing-basics
```

## 3. got / want

```go
got := ...
want := ...
```

- `got`：实际结果
- `want`：期望结果

测试核心：

```text
执行代码
↓
得到 got
↓
与 want 比较
↓
PASS / FAIL
```

## 4. `t.Errorf` 与 `t.Fatalf`

`t.Errorf`：

```text
标记测试失败
+
继续执行当前测试函数
```

`t.Fatalf`：

```text
标记测试失败
+
立即终止当前测试函数
```

如果当前失败已经导致后续断言无法安全继续，例如 `got == nil` 后还要访问 `got.ID`，应使用 `t.Fatal` / `t.Fatalf`。

## 5. Service 与 Repository

```go
type Repository interface {
	FindByID(id int64) (*User, error)
}

type Service struct {
	repo Repository
}
```

Service 依赖接口，因此测试时不需要真实数据库。

## 6. Fake Repository

```go
type fakeRepository struct {
	user *User
	err  error
}

func (r *fakeRepository) FindByID(id int64) (*User, error) {
	return r.user, r.err
}
```

Fake Repository 是可控的测试替身。

成功场景：

```text
user = Alice
err  = nil
```

失败场景：

```text
user = nil
err  = ErrUserNotFound
```

## 7. interface + DI 为什么利于测试

生产：

```text
Service
   ↓
Repository
   ↓
PostgresRepository
```

测试：

```text
Service
   ↓
Repository
   ↓
fakeRepository
```

Service 无需修改。

## 8. 成功路径测试

成功 contract：

```text
Repository 返回 User, nil
↓
Service 返回 User, nil
```

至少验证：

```text
err == nil
got != nil
got.ID 正确
got.Name 正确
```

## 9. 失败路径测试

Service 包装错误：

```go
return nil, fmt.Errorf("get user %d: %w", id, err)
```

失败 contract：

```text
got == nil
err != nil
errors.Is(err, ErrUserNotFound) == true
```

## 10. 为什么使用 `errors.Is`

Service 返回的不是原始 sentinel error 本身，而是包装后的 error。

所以不应该只依赖：

```go
err == ErrUserNotFound
```

而应该：

```go
errors.Is(err, ErrUserNotFound)
```

`errors.Is` 会沿 `%w` 保留的 error chain 检查根因。

```text
get user 999
     ↓
ErrUserNotFound
```

## 11. 当前测试状态

已完成：

```text
TestAdd
TestServiceGetUserSuccess
TestServiceGetUserNotFound
```

三个测试均通过。

## 12. Java / Python 对照

Java / JUnit：

```java
assertEquals(expected, actual);
```

Python / pytest：

```python
assert actual == expected
```

Go 标准 testing：

```go
if got != want {
	t.Errorf("got %v, want %v", got, want)
}
```

Go 标准测试更倾向直接使用普通 Go 控制流。

## 13. 本节核心结论

1. `_test.go` 是 Go 测试文件；
2. `TestXxx(t *testing.T)` 是测试函数；
3. `got` 是实际值，`want` 是期望值；
4. `t.Errorf` 报告失败但继续；
5. `t.Fatalf` 报告失败并立即结束当前测试；
6. interface + DI 允许测试时替换真实依赖；
7. Fake Repository 可以人为控制成功/失败路径；
8. `errors.Is` 用于识别包装后的错误链。
