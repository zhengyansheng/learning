package gc

import (
	"runtime"
	"testing"
)

func TestForceGC(t *testing.T) {

	m := make(map[string]string, 10)
	m["hello"] = "world"

	runtime.GC()
	t.Log("force gc")

	var forcegcperiod int64 = 2 * 60 * 1e9
	t.Log(forcegcperiod)
}
