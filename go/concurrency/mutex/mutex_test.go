package mutex

import (
	"sync"
	"testing"
	"time"
)

func TestMutexTryLock(t *testing.T) {

	var l sync.Mutex

	go func() {
		l.Lock()
		time.Sleep(time.Second * 3)
		l.Unlock()

	}()
	time.Sleep(time.Second)

	if l.TryLock() {
		t.Log("try lock success")
		defer l.Unlock()
	} else {
		t.Log("try lock failed")
	}
	t.Log("Done")
}
