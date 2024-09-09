package gc

import (
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"testing"
	"time"
)

func TestForceGC(t *testing.T) {

	m := make(map[string]string, 10)
	m["hello"] = "world"

	runtime.GC()
	t.Log("force gc")

	var forcegcperiod int64 = 2 * 60 * 1e9
	t.Log(forcegcperiod)
}

func TestGcTrace(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1<<20)
	}
	t.Logf("interval: %v", time.Since(start))
}

func TestSlice(t *testing.T) {
	start := time.Now()
	//var arr []int
	arr := make([]int, 1000)
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}
	t.Logf("interval: %v", time.Since(start))
}

// TestTraceFile go tool trace
func TestTraceFile(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		t.Failed()
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1<<20)
	}

}
func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	m := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			log.Printf("gc总次数: %d, 上次gc时间: %v, pauseTotal: %v", s.NumGC, s.LastGC, s.PauseTotal)
			runtime.ReadMemStats(&m)
			log.Printf("gc总次数: %d, 上次gc时间: %v, heap_obj_num: %v", m.NumGC, time.Unix(int64(time.Duration(m.LastGC).Seconds()), 0), m.HeapObjects)
		}
	}
}

func TestSlice2(t *testing.T) {
	go printGCStats()
	start := time.Now()
	var arr []int
	for i := 0; i < 10000000000; i++ {
		arr = append(arr, i)
	}
	t.Logf("interval: %v", time.Since(start))
}
