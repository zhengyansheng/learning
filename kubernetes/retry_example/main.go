package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/retry"
)

func main() {
	ch := make(chan int)
	go Do(ch)
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		// 如果匿名函数返回nil，则退出；如果匿名函数返回errors.NewConflict 则继续重试
		v, ok := <-ch
		if !ok {
			return nil
		}
		fmt.Printf("hello %d\n", v)
		return errors.NewConflict(schema.GroupResource{Resource: "test"}, "other", nil)
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}

func Do(ch chan int) {
	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}
