# unit05-testing

从"会写测试函数"到"会组织测试"：先用 `testing` 包跑通单个断言，再用测试表把同类用例批量驱动，最后给 fake 装上"黑匣子"验证协作者交互。核心 API：`t.Errorf`/`t.Fatalf`、`errors.Is`、`[]struct` + `for range` + `t.Run`、`reflect.DeepEqual`、参数/调用次数断言。

本 unit 是**本人手写练习**：目录内 notes.md 记题目与要点，代码先留空壳（如 `// 应该返回什么？`、`{name, ...}` 半截表），再由自己填全并 Review 校对。

## 知识点

| 目录 | 状态 | 一句话结论 |
|---|---|---|
| [01-testing-basics](01-testing-basics/)（2026-09-21） | ✅ | 测试文件 `xxx_test.go` + `func TestXxx(t *testing.T)` 才能被 `go test` 发现；`Errorf` 记失败继续跑、`Fatal(f)` 记失败立即终止——后续断言依赖当前结果时（如 `got.ID` 会 nil panic）必须 Fatal；Arrange-Act-Assert 三段式；fake 实现 `Repository` 接口即可替换真库；`Fatalf` 选型的判据是"失败了往下走还有没有意义" |
| [02-table-driven-tests](02-table-driven-tests/)（2026-09-21） | ✅ | 匿名 `[]struct` 表管数据、`for range` 管流程、`t.Run(tt.name, ...)` 管隔离与命名（子测试各拿自己的 `t`，`Fatalf` 只杀本 case）；错误判有无用 `(err != nil) != (wantErr != nil)`、判种类必须 `errors.Is`（`%w` 包装后 `==` 必挂）；`*User` 用 `==` 比的是地址，内容比较要 `reflect.DeepEqual`（含 slice/map 字段的 struct 用 `==` 会编译失败，DeepEqual 通用）；Go 1.22 起 range 变量每轮新建，闭包捕获不再需要 `tt := tt` |
| [03-interaction-testing](03-interaction-testing/)（2026-09-23） | ✅ | fake 加 `calls`/`lastID` 记录字段后，能抓住两个"返回值全对、行为却错"的 bug（传错 id、重复调用）——两个变异实验里结果断言全部通过，只有交互断言报红；返回值测试测 what（对重构不敏感），交互测试测 how（把测试绑在实现上，慎用） |

## 测试要点速查

- **测试发现规则是签名级的**：文件名必须 `_test.go` 结尾、函数必须 `Test` 前缀 + 大写开头 + 唯一参数 `*testing.T`，少一条 `go test` 静默跳过（不是通过）。
- **Errorf vs Fatalf 的判据**：当前断言失败会让后续断言"无意义或 panic"→ Fatal；各断言相互独立 → Errorf（一次跑完报出所有问题）。
- **`errors.Is` 而非 `==`**：service 里 `fmt.Errorf("...: %w", err)` 造出新错误对象，只有 `Is` 会沿 `Unwrap` 链找到哨兵错误。
- **表驱动三分工**：表 = 数据（repoUser/repoErr 是 fake 的输入，wantUser/wantErr 是期望输出），循环 = 流程，`t.Run` = 隔离；加 case 只加一行。
- **断言顺序也是分层**：先 err/got 结果（Fatal 系），后交互（Errorf 系）——结果都不对时查过程没意义。
- **Mock 的克制**：默认用真实对象测返回值；只有"过程本身是需求"（支付调两次 = 事故）或协作者昂贵/不可控（真库、HTTP、时钟）时才断言交互，且优先手写小 fake 而非 mock 框架。

## 与 unit04 的关系

unit04 把 `Repository` 接口定为包边界（`internal` 限制、接口按消费者设计）；unit05 立刻兑现这个设计的测试红利——正因为 service 依赖的是接口而非实现，测试里才能用一个 6 行的 `fakeRepository` 替换整个存储层，把"错误传播"（`ErrUserNotFound` + `%w` + `errors.Is`）这些 unit03/04 练过的链路在测试里闭环验证。
