package slice_append

import (
	"testing"
)

func TestSliceAppend(t *testing.T) {

	var slice []int
	var capSlice []int

	for i := 0; i < 2000; i++ {
		t.Logf("len: %v, cap: %v", len(slice), cap(slice))
		capSlice = append(capSlice, cap(slice))
		slice = append(slice, i)
	}

	newSlice := removeDuplicates(capSlice)
	t.Logf("切片的扩容机制: %v", newSlice)

}

func removeDuplicates(nums []int) []int {
	// 创建一个空的映射，用于跟踪已经出现的元素
	seen := make(map[int]struct{})
	result := []int{}

	for _, num := range nums {
		if _, exists := seen[num]; !exists {
			// 如果元素没有出现过，添加到结果切片和映射中
			seen[num] = struct{}{}
			result = append(result, num)
		}
	}
	return result
}
