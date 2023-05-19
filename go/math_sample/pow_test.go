package math_sample

import (
	"math"
	"testing"
	"time"
)

func TestMathPow(t *testing.T) {
	t.Log(math.Pow(2, float64(1)))
	t.Log(math.Pow(2, float64(2)))
	t.Log(math.Pow(2, float64(3)))
	t.Log(math.Pow(2, float64(4)))

	// 5*time.Millisecond, 1000*time.Second
	// backoff := float64(r.baseDelay.Nanoseconds()) * math.Pow(2, float64(exp))
	baseDelay := 5 * time.Millisecond
	backoff := float64(baseDelay.Nanoseconds()) * math.Pow(2, float64(1))
	t.Logf("backoff: %v, maxInt64: %v", backoff, math.MaxInt64)
	//t.Log(backoff < math.MaxInt64)
	t.Log(time.Duration(backoff))

	//maxDelay := time.Second * 1000

}
