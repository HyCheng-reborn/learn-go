# 运行验证

归档生成后已在 Go 1.23.2 环境中对每个含 `main.go` 的目录执行：

```bash
go run .
```

结果：12 / 12 通过。

已验证目录：

- unit01-basic/01-variables-and-error
- unit01-basic/02-control-flow-and-functions
- unit01-basic/03-string-and-rune
- unit02-collections/01-array-value-copy
- unit02-collections/02-slice-append-cap
- unit02-collections/03-map
- unit02-collections/04-pointer
- unit03-struct-method-interface/01-struct-and-constructor
- unit03-struct-method-interface/02-value-vs-pointer-receiver
- unit03-struct-method-interface/03-interface-method-set
- unit03-struct-method-interface/04-embedding-composition
- unit03-struct-method-interface/05-repository-service-di-IN_PROGRESS

注意：map 的遍历顺序本来就不应依赖；本归档示例没有依赖 map 的固定遍历顺序。
