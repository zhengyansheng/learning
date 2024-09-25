package recover

import "testing"

func TestRecover(t *testing.T) {

	// recover 函数没有直接被defer函数直接调用
	t.Run("recover_success_example", func(t *testing.T) {
		defer func() {
			if ok := recoverPanic(); !ok {
				t.Log("recover ...")
			}
		}()

		var n int = 0
		updateTable(n)

		t.Log("Finish")
	})

	t.Run("recover_success_example", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Logf("recover err: %v ...", err)
			}
		}()

		var n int = 0
		updateTable(n)

		t.Log("Finish")
	})

}

func updateTable(n int) {
	panic("err: n = 0")
}

func recoverPanic() bool {
	if err := recover(); err != nil {
		return false
	}
	return true
}
