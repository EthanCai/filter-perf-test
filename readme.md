# 介绍

这个例子用于比较`robpike.io/filter`和loop遍历数组的速度

These codes are used to compare the speed of `robpike.io/filter` and loop.

# 运行单元测试

```bash
# 运行全部单元测试
$ go test -v .

# 运行以GetIds1为前缀的单元测试
$ go test -run "GetIds.*" -v .
```

# 运行性能测试

```bash
# 运行全部性能测试
$ go test -v -bench=".*"

# 运行以GetIds为前缀的性能测试
$ go test -v -bench="GetIds.*" -count 1

# 生成性能测试报告
$ go test -v -bench="GetIds.*" -count 1 -cpuprofile=cpu.prof

# 使用go tool pprof工具查看性能测试报告
$ go tool pprof filter_perf_test.test cpu.prof
```

## 性能测试结果

My Computer's configuration:

    MacBook Pro (Retina, 15-inch, Mid 2014)
    OS: OS X EI Capitan Version 10.11.3
    Processor: 2.8 GHz Intel Core i7
    Memory: 16 GB 1600 MHz DDR3

Test Result:

```bash
$ go version

go version go1.6 darwin/amd64

$ go test -v -bench="GetIds.*" -count 1

=== RUN   TestCreateElements
--- PASS: TestCreateElements (0.00s)
=== RUN   TestGetIds1
--- PASS: TestGetIds1 (0.00s)
=== RUN   TestGetIds2
--- PASS: TestGetIds2 (0.00s)
PASS
BenchmarkGetIds1-8	2016/03/09 18:13:27 b.N = 1
2016/03/09 18:13:27 b.N = 100
2016/03/09 18:13:27 b.N = 10000
2016/03/09 18:13:27 b.N = 1000000
2016/03/09 18:13:28 b.N = 5000000
 5000000	       375 ns/op
BenchmarkGetIds2-8	2016/03/09 18:13:31 b.N = 1
2016/03/09 18:13:31 b.N = 100
2016/03/09 18:13:31 b.N = 10000
2016/03/09 18:13:31 b.N = 1000000
2016/03/09 18:13:31 b.N = 100000000
2016/03/09 18:13:57 b.N = 300000000
300000000	        26.2 ns/op
ok  	github.com/EthanCai/filter_perf_test	116.002s
```

`375 ns / 26.2 ns = 14.31`

# 相关知识

- [Profiling go programs](http://blog.golang.org/profiling-go-programs)
- [How to write benchmarks in go](http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [go test](http://docs.studygolang.com/cmd/go/#hdr-Test_packages)
- [description of testing flags](http://docs.studygolang.com/cmd/go/#hdr-Description_of_testing_flags)
- [testing package](https://golang.org/pkg/testing/)
