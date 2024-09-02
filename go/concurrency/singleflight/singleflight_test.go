package singleflight

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

/*
并发请求中，用singleFlight可以避免缓存击穿，减少对数据库的压力
*/
func TestSingleFlight(t *testing.T) {
	testCases := []struct {
		name       string
		concurrent int
		expected   string
	}{
		{"single flight", 1000, "article: 1"},
	}

	for _, tc := range testCases {
		result, err := singleFlightFunc(tc.concurrent)
		if err != nil {
			t.Fatal(err)
		}
		if result != tc.expected {
			t.Errorf("Name: %v, Expected %v, but got %v", tc.name, tc.expected, result)
		}
	}
}

func singleFlightFunc(concurrent int) (string, error) {
	var (
		count int32
		wg    sync.WaitGroup
		now   = time.Now()
		s     = &singleflight.Group{}
	)

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 模拟并发请求
			res, _ := getArticleForSingleFlight(s, count, 1)
			if res != "article: 1" {
				panic("wrong")
			}
		}()
	}
	wg.Wait()
	fmt.Printf("concurrent: %d, time: %v\n", concurrent, time.Since(now))
	return "article: 1", nil
}

func getArticleForSingleFlight(s *singleflight.Group, count int32, i int) (string, error) {
	v, err, _ := s.Do(fmt.Sprintf("article-%d", i), func() (interface{}, error) {
		return getArticle(count, i)
	})
	return v.(string), err
}

func getArticle(count int32, i int) (string, error) {
	// 假设这里会对数据进行调用，模拟不同并发下耗时不同
	atomic.AddInt32(&count, 1) // 原子操作+1
	time.Sleep(time.Duration(count) * time.Millisecond)

	return fmt.Sprintf("article: %d", i), nil
}
