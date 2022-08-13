#### `go test` 工具

`go test`命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以 `_test.go` 的文件都会被测试，不会被`go build`编译到最终的可执行文件中。

```shell
# 不加参数，到文件级别
go test

PS D:\git\go\go_demo\test\testDemo> go test
PASS
ok      go_demo/test/testDemo   0.040s

# -v 添加函数名称和执行时间
go test -v

PS D:\git\go\go_demo\test\testDemo> go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSquare
--- PASS: TestSquare (0.00s)
PASS
ok      go_demo/test/testDemo   0.025s

# -run 只测试指定的函数，正则匹配
PS D:\git\go\go_demo\test\testDemo> go test -v -run "Squ"
=== RUN   TestSquare
--- PASS: TestSquare (0.00s)
PASS
ok      go_demo/test/testDemo   0.032s

# cover 覆盖率
# coverprofile=c.out
PS D:\git\go\go_demo\test\testDemo> go test -cover  
PASS
coverage: 100.0% of statements
ok      go_demo/test/testDemo   0.040s

# bench 基准测试
# -bench=. 为全部可测的函数，也可以 -bench=Add 只测试指定函数
# -benchmem 内存信息
# -benchtime=20s 最小基准时间，默认为1秒
D:\git\go\go_demo\test\testDemo>go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: go_demo/test/testDemo
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkAdd-8          1000000000               0.2562 ns/op          0 B/op          0 allocs/op
BenchmarkSquare-8       1000000000               0.2589 ns/op          0 B/op          0 allocs/op
PASS
ok      go_demo/test/testDemo   0.616s
```

#### 函数类型

> 函数名除了前缀，首字母要大写。

| 类型    | 前缀        | 参数           | 说明              |
|-------|-----------|--------------|-----------------|
| 测试函数  | Test      | t *testing.T | 逻辑测试            |
| 基准测试  | Benchmark | b *testing.B | 性能测试            |
| 示例函数  | Example   | 无            | 为文档提供示例文档       |
