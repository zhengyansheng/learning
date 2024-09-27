package unit_test

import (
	"testing"
	"unit_test/funcs"
)

// TestAdd 单元测试
func TestAdd(t *testing.T) {
	var x int = 2
	var y int = 3
	var expected = 5 // 预期的

	actual := funcs.Add(x, y) // actual 实际的
	if actual != expected {
		t.Errorf("Add(%d %d) = %d; excepted: %d", x, y, actual, expected)
	}
}

// TestAdd 单元测试
func TestSubAdd(t *testing.T) {
	var x int = 2
	var y int = 3
	var expected = 5 // 预期的

	t.Run("sub add1", func(t *testing.T) {
		actual := funcs.Add(x, y) // actual 实际的
		if actual != expected {
			t.Errorf("Add(%d %d) = %d; excepted: %d", x, y, actual, expected)
		}
	})

	t.Run("sub add2", func(t *testing.T) {
		actual := funcs.Add(x, y) // actual 实际的
		if actual != expected {
			t.Errorf("Add(%d %d) = %d; excepted: %d", x, y, actual, expected)
		}
	})

	t.Run("sub add3", func(t *testing.T) {
		actual := funcs.Add(x, y) // actual 实际的
		if actual != expected {
			t.Errorf("Add(%d %d) = %d; excepted: %d", x, y, actual, expected)
		}
	})
}
