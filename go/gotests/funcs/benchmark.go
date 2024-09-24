package funcs

// MakeSliceWithoutAlloc 不预分配内存
func MakeSliceWithoutAlloc() []int {
	var newSlice []int

	for i := 0; i < 1000000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}

// MakeSliceWithAlloc 预分配内存
func MakeSliceWithAlloc() []int {
	var newSlice []int

	newSlice = make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}
