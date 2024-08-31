package value_pointer

import "testing"

func TestValuePointer(t *testing.T) {

	t.Run("int64", func(t *testing.T) {
		var i int64 = 1
		t.Logf("value: %v", i)
		modifyInt(i)
		t.Logf("value: %v", i)

	})

}

func modifyInt(i int64) {
	i = 100
}
