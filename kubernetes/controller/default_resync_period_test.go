package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/* cmd/kube-controller-manager/app/controllermanager.go
func ResyncPeriod(c *config.CompletedConfig) func() time.Duration {
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(c.ComponentConfig.Generic.MinResyncPeriod.Nanoseconds()) * factor)
	}
}
*/

func ResyncPeriod() func() time.Duration {
	return func() time.Duration {
		// [1.0,2.0)
		factor := rand.Float64() + 1
		fmt.Printf("factor: %v\n", factor)
		return time.Duration(float64(2) * factor)
	}
}

func TestResyncPeriod(t *testing.T) {
	d := ResyncPeriod()
	t.Logf("time duration: %v", d().String())
	time.Sleep(d())
	fmt.Println("stop")
}
