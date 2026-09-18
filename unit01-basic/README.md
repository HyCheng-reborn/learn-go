# unit01-basic

基础语法与错误处理。

## 知识点

| 目录 | 状态 | 一句话结论 |
|---|---|---|
| [01-variables-and-error](01-variables-and-error/) | ✅ 2026-09-15 亲手写 | 哨兵错误要用 `errors.Is` 判断；`if result, err := f(); err == nil` 能把变量限制在 if 作用域内；自定义的 `max` 会遮蔽 Go 内置的 `max` |
| [02-control-flow-and-functions](02-control-flow-and-functions/) | 📦 2026-09-18 归档导入（重构版） | `for` 三段式累加；`switch` 可以不写表达式，直接 `case n < 0:`；用 `(int, bool)` 多返回值表达失败，是 `error` 之外的另一种习惯写法 |
| [03-string-and-rune](03-string-and-rune/) | 📦 2026-09-18 归档导入（重构版） | `len(string)` 是**字节数**不是字符数："Go语言" 是 8 字节 / 4 rune；`range` 遍历 string 得到 rune，下标是字节位置且不连续（0、1、2、5） |

## 涉及内容

变量声明与类型推断、`Println` / `Printf` 格式动词、多返回值函数、`errors.New`、哨兵错误 `ErrUserNotFound`、`for` / `switch` / `if`、string 与 byte / rune 的区别。

## 关于 01 的 notes.md

`01-variables-and-error/notes.md` 来自归档，其中的代码片段是归档版 `findUser` 的写法，与你亲手写的 `main.go` 不完全相同（你的版本还包含 `max`、`isEven`、`divide` 和几段注释掉的练习）。**以 main.go 为准**，notes.md 只作概念说明。

notes.md 里有一条正好解释了你的 main.go 的行为：一旦进入 `if err != nil { return }`，`return` 之后的语句不会再执行 —— 所以 `findUser("Bob")` 之后那句 `fmt.Printf("user: %s\n", user)` 永远不会打印。

## 归档来源说明

`02`、`03` 在归档中被明确标注为「按当时知识点**重构**的复习例子，不是声称与当时文件逐字一致」。也就是说这两个 `main.go` 是整理时重写的，不是你当年的原始代码；那部分原始练习在早期覆盖中已经丢失，无法逐字还原。
