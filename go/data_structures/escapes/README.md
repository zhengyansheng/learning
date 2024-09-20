# 逃逸分析

## 执行
```bash
# go run -gcflags="-m -m -l" main.go
./main.go:8:2: res escapes to heap: // res逃逸到堆
./main.go:8:2:   flow: ~r0 = &res:
./main.go:8:2:     from &res (address-of) at ./main.go:10:9
./main.go:8:2:     from return &res (return) at ./main.go:10:2
./main.go:8:2: moved to heap: res  // 移动到堆：res
```

## 分析
```
一般来说，局部变量会在函数结束后被销毁。

如果局部变量的作用域超出了函数，则不会将内存分配在栈上，而是分配在堆上，因为他们不再栈区
即使释放函数，其内容也不受影响.

```


```bash
# go build -gcflags -m main.go
```
