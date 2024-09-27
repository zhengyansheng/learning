# Go

## [概念](./intra.md)
- [x] GMP并发模型
- [x] GC垃圾回收策略 (三色标记法+混合写屏障)
- [x] GC垃圾回收何时触发
- [ ] 内存逃逸

## 数据结构
- [ ] array 和 slice 的区别
- [x] [slice 的扩容机制](./data_structures/slice_expansion/README.md)
- [x] [map 和 sync.Map](./concurrency/syncmap/READNE.md)
- [x] [string/array/slice/map 是值传递还是引用传递](./data_structures/value_pointer/README.md)
- [x] [堆(heap) 栈(stack)的区别](./data_structures/heap_stack/README.md)
- [x] [channel是分配在堆上还是栈上](./data_structures/channel/README.md)
- [x] [make 和 new](./data_structures/make_new/README.md)
- [ ] channel底层的实现
- [ ] 内存管理

## [分布式](./distributed.md)
- [x] 消息队列中，如何处理重复消息
- [x] 日志的定位

## 设计模式
- [ ] 函数选项模式 （支持默认值）
- [ ] 装饰器模式