# Go 单元测试

通过 `-v` 显示执行结果，要不只能看到一个运行是否成功的提示和消耗时间

```bash
go test -v pkg/4_arrays_and_slices/array/*.go
```

添加 `-run` 以指定想要运行的测试函数

```bash
go test -run ^TestArrays$ -v pkg/4_arrays_and_slices/array/*.go
```
