# unit02-collections

集合与接口。`01`、`02` 是早期学习路径留下的历史目录（由原来一个 `main.go` 拆分而来），`03` 起为 2026-09-18 归档导入。

## 知识点

| 目录 | 状态 | 一句话结论 |
|---|---|---|
| [01-slice-append-cap](01-slice-append-cap/) | ✅ 2026-09-16 亲手写 | slice 赋值只复制 len/cap/底层数组指针，共享同一数组；`append` 超过 cap 时分配新底层数组，之后 `a` 与 `b` 彻底脱钩（本机实测 `a` 变 len:4 cap:6，`b` 仍是 `[10 20 30]`） |
| [02-interface-receivers](02-interface-receivers/) | ✅ 2026-09-17 亲手写 | method set 规则：**值接收者 → `T` 和 `*T` 都实现接口；指针接收者 → 只有 `*T` 实现接口，`T` 不实现**。但 `dog.Speak()` 这样的普通方法调用会自动取地址，所以值也能调，只有把值赋给接口变量才会编译失败 |
| [03-array-value-copy](03-array-value-copy/) | 📦 归档导入 | array 是**值类型**：`b := a` 复制整个数组，改 `b[0]` 不影响 `a`。这与 slice 的共享底层数组正好相反 |
| [04-map-comma-ok](04-map-comma-ok/) | 📦 归档导入（重构版） | `v, ok := m[k]` 用来区分「值是零值」和「键不存在」；`delete` 删不存在的键不报错；map 遍历顺序不保证，示例不依赖它 |
| [05-pointer-basics](05-pointer-basics/) | 📦 归档导入（重构版） | `&n` 取地址、`*p` 解引用；通过 `*p = 20` 能改到原变量；函数收 `*int` 才能修改调用方的值（`addOne(&n)`） |

## 关于 01 的 cap 数字

`01-slice-append-cap/notes.md`（归档导入）里有一条重要提醒，修正本表容易造成的误解：

> 不要把「cap 一定从 3 变成 6」当成语言规范。

具体扩容容量属于**运行时实现细节**。真正要掌握的是：`append` 超过当前容量时可能分配新底层数组，原本共享数据的 slice 从那一刻开始脱钩。上表的 `len:4 cap:6` 只是 go1.27.1 / windows-amd64 上的一次实测值。

## 涉及内容

`make` 创建 slice、len/cap、append 扩容与底层数组、array 值语义、map 增删改查与 comma-ok、指针 `&` / `*`、struct 定义与键值初始化、方法接收者（值 / 指针）、接口隐式实现、接口作参数实现多态。

## 待补

`range` 遍历 slice/map 的完整用法、slice 的三个下标切片 `a[low:high:max]`、`copy` 的语义。`for` / `switch` 已在 `unit01-basic/02-control-flow-and-functions` 覆盖。

## 关于 02 的位置

从知识分类看 interface 不属于 collections，`02-interface-receivers` 留在这里是历史学习路径造成的。本轮不移动、不重命名；`unit03-struct-method-interface` 里有同主题的归档整理版，两者并存。
