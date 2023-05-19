package time_after

import (
	"context"
	"testing"
	"time"

	"k8s.io/utils/clock"
)

func TestTimeAfter(t *testing.T) {
	/*
		After 大于 「now」 时间
		Before 小于 「now」 时间
	*/
	readyAt, _ := time.Parse("2006-01-02 15:04:05", "2023-04-19 19:00:00")
	now := time.Now()
	t.Logf("readyAt(%v) < now(%v), %v", readyAt, now, readyAt.After(now))
	t.Logf("readyAt(%v) < now(%v), %v", readyAt, now, readyAt.Before(now))

	readyAt2, _ := time.Parse("2006-01-02 15:04:05", "2023-06-19 19:00:00")
	t.Logf("readyAt2(%v) > now(%v), %v", readyAt2, now, readyAt2.After(now))
	t.Logf("sub %v", readyAt2.Sub(now))
}

func TestTimeC(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	var ck clock.Clock
	ck = clock.RealClock{}
	t.Log(ck.Now())
	timer := ck.NewTimer(time.Second * 5)
	nextReadyAt := timer.C()

	select {
	case <-ctx.Done():
		t.Log("ctx done")
	case <-nextReadyAt:
		t.Log("timer done")
	}

}
