package ticker

import (
	"testing"
	"time"
)

/*
	NewTimer 一次性定时器，仅会执行一次
	NewTicker 周期性定时器，周期性执行
*/

func TestTimer(t *testing.T) {

	t.Run("timer stop", func(t *testing.T) {
		timer := time.NewTimer(time.Second * 3)
		defer timer.Stop() // 关闭Timer 否则会引起内存泄漏

		<-timer.C

		t.Log("timer done")
	})

	t.Run("timer stop", func(t *testing.T) {
		timer := time.NewTimer(time.Second * 3)
		defer timer.Stop() // 关闭Timer 否则会引起内存泄漏

		select {
		case <-timer.C:

		}

		t.Log("timer done")
	})

	t.Run("timer reset", func(t *testing.T) {

	})

	t.Run("timer after", func(t *testing.T) {
		<-time.After(time.Second * 3)
		t.Log("hello world")
	})

	t.Run("timer afterFunc 异步执行", func(t *testing.T) {
		time.AfterFunc(time.Second*2, func() {
			t.Log("hello world")
		})

		time.Sleep(time.Second * 4)
	})

	t.Run("ticker", func(t *testing.T) {
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop() // 关闭Timer 否则会引起内存泄漏

		for {
			select {
			case <-ticker.C:
				t.Log("didi didi ......")
			}
		}
	})

}
