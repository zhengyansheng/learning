package wait_util

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func TestPollImmediateUntil(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := wait.PollImmediateUntil(time.Second*1, func() (bool, error) {
		t.Log("continue")
		return true, nil
		// return true, fmt.Errorf("xxx")
	}, ctx.Done())
	if err != nil {
		panic(err)
	}
}
