package array_slice

import (
	"fmt"
	"testing"
	"time"
)

// 1. 数组和切片的定义和使用
func TestDefineArraySlice(t *testing.T) {
	fmt.Println("[数组]====================================")
	defineArray()
	fmt.Println("[切片1]====================================")
	defineSlice1()
	fmt.Println("[切片2]====================================")
	defineSlice2()
}

// 2. 数组是值类型，切片是引用类型
func TestModifyArraySliceTest(t *testing.T) {
	fmt.Println("[修改]====================================")

	var arr = [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Printf("arr: %v\n", arr)

	var slice = []int{1, 2, 3}
	modifySlice(slice)
	fmt.Printf("slice: %v\n", slice)

}

// 3. 切片扩容
func TestScaleSlice(t *testing.T) {
	scaleSlice()
}

func defineArray() {
	var arr [3]int
	fmt.Printf("cap: %d, arr: %v\n", cap(arr), arr)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	//arr[3] = 4
	fmt.Printf("cap: %d, arr: %v\n", cap(arr), arr)
}

func defineSlice1() {
	var slice []int
	fmt.Printf("cap: %d, slice: %v\n", cap(slice), slice)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	fmt.Printf("cap: %d, slice: %v\n", cap(slice), slice)
}

func defineSlice2() {
	slice := make([]int, 3)
	fmt.Printf("cap: %d, slice: %v\n", cap(slice), slice)
	slice[0] = 5
	slice[1] = 6
	slice[2] = 7
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)

	fmt.Printf("cap: %d, slice: %v\n", cap(slice), slice)
}

// modifyArray 数组是值类型，传递的是数组的副本
func modifyArray(arr [3]int) {
	arr[0] = 100
}

// modifySlice 切片是引用类型，传递的是切片的地址
func modifySlice(slice []int) {
	slice[0] = 100
}

func scaleSlice() {
	var slice []int
	// 1,2,4,8,16,32,64,128,256,512,848,1280
	// 如果小于1024，每次扩容2倍
	// 如果大于1024，每次扩容1.5倍

	for i := 0; i < 1025; i++ {
		slice = append(slice, i)
		fmt.Printf("cap: %d, len: %d, slice: %v\n", cap(slice), len(slice), slice)
		time.Sleep(time.Millisecond * 20)
	}
}
