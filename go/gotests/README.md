# 测试

## 单元测试
> testing.T


## 基准/性能测试
> testing.B


## 模糊测试
> testing.F

### 例子
```go
func FuzzReverse(f *testing.F) {
	...
	f.Fuzz(func(t *testing.T, orig string) {
		...
	}
}
```

### 运行
```bash
# go test -fuzz=FuzzReverse . -fuzztime=10s
```

## 子测试

### 例子

```go

func TestSubExample(t *testing.T) {
	
    t.Run("sub example1", func (t *testing.T) {
        ...
    })

    t.Run("sub example2", func (t *testing.T) {
        ...
    })
    
    t.Run("sub example3", func (t *testing.T) {
        ...
    })
}
```

### 运行
```bash
# go test -v -run TestSubAdd/sub_add1 .
```