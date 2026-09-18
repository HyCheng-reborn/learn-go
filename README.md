# learn-go

Go 语言学习项目。每个知识点是一个独立目录，内含一个可直接运行的 `package main` 示例。

## 约定

- **一个知识点 = 一个目录 = 一个 `main.go`**（同一目录下只能有一个 `func main`，所以新知识点必须开新目录）。
- 目录**只新增，不改写**：学新东西时不覆盖旧示例，旧代码长期保留。
- unit 目录用 `unitNN-主题`，知识点目录用 `NN-要点`，unit 内编号从 `01-` 重新开始，全部小写 kebab-case。
- 同目录可放 `notes.md`，记录核心结论、易混点、Java / Python 对照。
- 历史练习代码优先级最高：整理版资料**不得**覆盖亲手写的示例。

## 单元清单

| 单元 | 状态 | 知识点 |
|---|---|---|
| [unit01-basic](unit01-basic/README.md) | ✅ | 变量与基本类型、错误处理（哨兵错误 + `errors.Is`）、控制流与函数、string / byte / rune |
| [unit02-collections](unit02-collections/README.md) | ✅ | slice 底层数组与 append 扩容、接口与方法接收者、array 值复制、map 与 comma-ok、指针基础 |
| [unit03-struct-method-interface](unit03-struct-method-interface/README.md) | 🚧 | struct 与构造函数式函数、值 / 指针接收者、interface 与 method set、embedding 组合、repository→service 依赖注入（进行中） |

每个 unit 的 README 里有逐条知识点表和「一句话结论」。

## 运行

```bash
# 从项目根目录运行任意一个知识点
go run ./unit01-basic/01-variables-and-error
go run ./unit02-collections/03-array-value-copy
go run ./unit03-struct-method-interface/05-repository-service-di

# 一次性检查全部示例
go vet ./...
gofmt -l .
go test ./...
```

在编辑器中也可以直接点 `func main` 上方的 CodeLens `run | debug`，无需记路径。

> Windows 上偶发 `go run` 报 `unlinkat ...\exe\xxx.exe: Access is denied` —— 这是杀毒软件短暂锁住刚生成的临时 exe 导致清理失败，程序本身已经正常输出完毕，不是代码问题。可用 `go build -o` 到别处再执行来绕开。

## notes/

不是可运行示例，是资料与快照：

| 文件 | 内容 |
|---|---|
| [env-go1.27.1.md](notes/env-go1.27.1.md) | 本机 Go 工具链快照（已脱敏：完整 PATH 与用户名已移除） |
| [go-learning-progress-20260918.md](notes/go-learning-progress-20260918.md) | 归档时的学习进度：已掌握清单 + 当前正在学的 interface/DI |
| [archive-validation-20260918.md](notes/archive-validation-20260918.md) | 归档生成时在 Go 1.23.2 下的 12/12 运行验证记录 |

## 资料来源：2026-09-18 归档导入

`unit01` 的 `02`/`03`、`unit02` 的 `03`~`05`、整个 `unit03` 来自 ChatGPT 整理的 `go_learning_archive_20260918.zip`（25388 B，sha256 `591aeea8…212c77f5`，导入前校验一致）。导入时的处理：

- **未采用 ZIP 自带的 `go.mod`**（`module learn-go` / `go 1.23.0`），保留本仓库根 `go.mod`（`go 1.27.1`），避免出现第二个根 module。
- **两个重名知识点保留你的历史版本**，只并入 `notes.md`：`unit01-basic/01-variables-and-error`、`unit02-collections/01-slice-append-cap`。ZIP 里的 `main.go` 未落地。
- **ZIP 顶层 `README.md` 未整文件导入**，其课程目录与「恢复说明」合并进本文件和各 unit README。`LEARNING_PROGRESS.md`、`VALIDATION.md` 原样存入 `notes/`。
- **目录重新编号**：ZIP 的 `unit02-collections/01-array-value-copy`、`03-map`、`04-pointer` 在本仓库是 `03-array-value-copy`、`04-map-comma-ok`、`05-pointer-basics`，因为 `01`、`02` 已被历史目录占用。
- **改名**：`05-repository-service-di-IN_PROGRESS` → `05-repository-service-di`（大写后缀不符合命名规则），进行中状态改记在 unit03 README。
- **可信度标注**：归档自己声明 `control flow / function`、`string / rune`、`map`、`standalone pointer` 这四项是「按当时知识点**重构**的复习例子，不是与当时文件逐字一致」。对应本仓库的 `unit01-basic/02`、`unit01-basic/03`、`unit02-collections/04`、`unit02-collections/05`，在各 unit README 中标为「重构版」。其余示例归档声称为可较可靠还原的版本。
- **本机复验**：归档记录的是 Go 1.23.2 下 12/12 通过；在本机 go1.27.1 / windows-amd64 重新逐个 `go run .`，12 个示例输出全部正确。

尚未学：泛型、goroutine / channel / 并发、context、testing、HTTP / Gin / Eino 等框架。相关目录一律不提前创建。
