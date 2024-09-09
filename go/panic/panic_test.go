package panic

import (
	"log"
	"sync"
	"testing"
)

// TestSystemPanic 系统级别的panic 拦截不住
func TestSystemPanic(t *testing.T) {

	var wg sync.WaitGroup
	var m = make(map[int]int)
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		for {
			t.Log(m[0])
		}
	}()

	wg.Wait()
	t.Log("Done")

}

// TestUserPanic 用户级别的panic 可以被拦截
func TestUserPanic(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			process(i)
		}()
	}

	wg.Wait()
	t.Log("Done")
}

func process(i int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if i == 0 {
		panic("0 is err")
	}
	log.Println(i)
}
