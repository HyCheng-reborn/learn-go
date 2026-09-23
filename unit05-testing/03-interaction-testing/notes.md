# Unit 05 · 03 Interaction Testing 学习笔记

## 1. 本节目标

本节学习测试中的“交互行为验证”。

之前的测试主要验证：

```text
输入
↓
Service
↓
输出结果是否正确
```

这一节进一步验证：

```text
Service 是否正确调用依赖？
调用了几次？
传入了什么参数？
```

---

## 2. 可记录交互的测试替身

```go
type fakeRepository struct {
    user *User
    err  error

    calls  int
    lastID int64
}
```

其中：

```text
user / err
→ 控制 Repository 返回什么

calls
→ 记录 FindByID 被调用次数

lastID
→ 记录最后一次收到的 id
```

实现：

```go
func (r *fakeRepository) FindByID(id int64) (*User, error) {
    r.calls++
    r.lastID = id

    return r.user, r.err
}
```

---

## 3. 输出行为与交互行为

### 输出行为

验证：

```text
GetUser 返回的 user 是否正确？
error 是否正确？
```

### 交互行为

验证：

```text
Repository 是否被调用？
调用了几次？
参数是否正确？
```

例如：

```go
if repo.calls != 1 {
    t.Errorf(...)
}

if repo.lastID != 42 {
    t.Errorf(...)
}
```

---

## 4. 为什么只检查输出可能漏掉 bug

假设 Service 错误地调用：

```go
repo.FindByID(999)
```

而 Fake 完全忽略 id，只返回预设的 Alice。

最终结果仍可能正确，因此只看输出的测试可能 PASS。

如果测试同时检查：

```go
repo.lastID
```

就能发现参数错误。

---

## 5. 为什么调用次数有时也重要

假设 Service 写成：

```go
s.repo.FindByID(id)
usr, err := s.repo.FindByID(id)
```

最终仍可能返回正确用户，但真实 Repository 被调用了两次。

如果真实实现查询数据库，就可能产生重复 SQL。

因此可以验证：

```go
repo.calls == 1
```

---

## 6. Fake / Stub / Mock 的实用理解

当前阶段不需要死记严格术语，可以先这样理解：

```text
Stub / 简单 Fake
→ 控制依赖返回什么

Mock-style interaction test
→ 还验证依赖是怎么被调用的
```

当前 `fakeRepository` 同时承担：

```text
控制返回值
+
记录调用行为
```

因此已经带有明显的 mock-style interaction verification。

---

## 7. 什么时候值得验证交互

交互本身具有业务或工程意义时值得验证，例如：

```text
不能重复扣款
不能重复发送邮件
不能重复调用昂贵的 LLM
必须传递正确的 userID
必须使用正确的 prompt / 参数
```

---

## 8. 不要过度 Mock

并不是所有内部调用都值得验证。

如果测试要求：

```text
必须先调用 A
再调用 B
再调用 C
每个调用必须恰好一次
```

而这些顺序只是内部实现细节，那么正常重构也可能让测试大量失败。

因此原则是：

> 如果只关心最终可观察结果，就优先测试结果。

> 如果交互本身就是重要 contract，再验证调用次数、参数或顺序。

---

## 9. 与 AI 后端的关系

未来测试 LLM Service 时，不希望每次测试都真实请求模型并消耗 Token。

可以使用测试替身：

```text
fakeLLMClient
```

它既可以返回固定 LLM 响应，又可以记录：

```text
调用次数
Prompt
模型参数
请求内容
```

因此当前的：

```text
fakeRepository
calls
lastID
```

以后可以直接迁移为：

```text
fakeLLMClient
calls
lastPrompt
```

思想完全一致。

---

## 10. 本节核心结论

1. 结果测试验证“最后得到什么”；
2. 交互测试验证“依赖是怎么被调用的”；
3. 简单 Fake 可以控制返回值；
4. Mock-style 测试会进一步验证调用次数、参数、顺序等；
5. 只有当交互本身有业务/工程意义时，才应该重点验证；
6. 不应为了使用 Mock 而过度绑定内部实现细节。
