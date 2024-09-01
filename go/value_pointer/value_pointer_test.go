package value_pointer

import (
	"testing"
)

func TestValuePointer(t *testing.T) {

	t.Run("int64", func(t *testing.T) {
		var i int64 = 1
		t.Logf("value: %v", i)
		modifyInt(i)
		t.Logf("value: %v", i)

		var ii int64 = 1
		t.Logf("value: %v", ii)
		modifyIntPointer(&ii)
		t.Logf("value: %v", ii)
	})

	t.Run("slice", func(t *testing.T) {
		var i []string = []string{"1", "2"}
		t.Logf("value: %v", i)
		modifySlice(i)
		t.Logf("value: %v", i)

		var ii []string = []string{"1", "2"}
		t.Logf("value: %v", ii)
		modifySlicePointer(&ii)
		t.Logf("value: %v", ii)
	})

	t.Run("map", func(t *testing.T) {
		x := make(map[string]string)
		x["1"] = "hello"
		x["2"] = "world"
		t.Logf("value: %v", x)
		modifyMap(x)
		t.Logf("value: %v", x)

		var y = make(map[string]string)
		y["1"] = "hello"
		y["2"] = "world"
		t.Logf("value: %v", y)
		modifyMapPointer(&y)
		t.Logf("value: %v", y)
	})

}

func modifyInt(i int64) {
	i = 100
}

func modifyIntPointer(i *int64) {
	*i = 100
}

func modifySlice(i []string) {
	//i[1] = "hello"
	i = append(i, "100")
}

func modifySlicePointer(i *[]string) {
	(*i)[0] = "100"
}

func modifyMap(i map[string]string) {
	i["3"] = "nihao"
	i["4"] = "java"
}

func modifyMapPointer(i *map[string]string) {
	(*i)["3"] = "python"
}
