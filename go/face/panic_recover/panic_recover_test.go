package panic_recover

import (
	"fmt"
	"testing"
	"time"
)

func TestPanicRecover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	// 引发 panic
	panic("Some error occurred")
}

func TestPanicRecover2(t *testing.T) {
	go func() {
		//defer func() {
		//	if r := recover(); r != nil {
		//		fmt.Println("Recovered:", r)
		//	}
		//}()

		// 引发 panic
		panic("Some error occurred")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("end")
}
