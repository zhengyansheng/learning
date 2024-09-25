# recover

_要满足三个条件_
- panic 时指定的参数为nil，比如 panic("xxx failed ...")
- 当前协程没有发生 panic
- recover 没有被defer函数直接调用