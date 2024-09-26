# channel是分配在堆上还是栈上

- channel默认是分配在栈上的
- 如果超出了函数的作用域则会逃逸到堆上，比如函数返回一个channel就会逃逸

```bash
# go build -gcflags="-m -m -l" .\main.go
./main.go:18:5: func literal escapes to heap:
./main.go:18:5:   flow: {heap} = &{storage for func literal}:
./main.go:18:5:     from func literal (spill) at ./main.go:18:5
./main.go:16:2: producer capturing by value: ch (addr=false assign=false width=8)
./main.go:18:5: func literal escapes to heap
./main.go:11:14: <-ch escapes to heap:
./main.go:11:14:   flow: {storage for ... argument} = &{storage for <-ch}:
./main.go:11:14:     from <-ch (spill) at ./main.go:11:14
./main.go:11:14:     from ... argument (slice-literal-element) at ./main.go:11:13
./main.go:11:14:   flow: {temp} = &{storage for ... argument}:
./main.go:11:14:     from ... argument (spill) at ./main.go:11:13
./main.go:11:14:     from log.Println(... argument...) (call parameter) at ./main.go:11:13
./main.go:11:14:   flow: {heap} = *{temp}:
./main.go:11:13: ... argument does not escape
./main.go:11:14: <-ch escapes to heap  // 这里发生逃逸到堆上
```