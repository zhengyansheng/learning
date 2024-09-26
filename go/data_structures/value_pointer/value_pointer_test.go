package value_pointer

import (
	"testing"
)

func TestValuePointer(t *testing.T) {

	t.Run("string", func(t *testing.T) {
		data := "hello"
		t.Logf("data: %v", data)
		modifyString(data)
		t.Logf("data: %v", data)
	})

	t.Run("array", func(t *testing.T) {
		data := [2]string{"1", "2"}
		t.Logf("data: %v", data)
		modifyArray(data)
		t.Logf("data: %v", data)
	})

	t.Run("slice", func(t *testing.T) {
		data := []string{"1", "2"}
		t.Logf("data: %v", data)
		modifySlice(data)
		t.Logf("data: %v", data)
	})

	t.Run("map", func(t *testing.T) {
		x := make(map[string]string)
		x["1"] = "hello"
		x["2"] = "world"
		t.Logf("value: %v", x)
		modifyMap(x)
		t.Logf("value: %v", x)
	})

}
func modifyString(i string) {
	i = "helloworld"
}

func modifyArray(i [2]string) {
	i[0] = "nihao"
}

func modifySlice(i []string) {
	//这会影响到原始切片，因为它们共享同一个底层数组。
	i[1] = "hello"

	// 这个修改不会影响原始切片，因为此时 切片结构体被复制，并指向一个新的底层数组。
	i = append(i, "100")
}

func modifyMap(i map[string]string) {
	i["3"] = "nihao"
	i["4"] = "java"
}
