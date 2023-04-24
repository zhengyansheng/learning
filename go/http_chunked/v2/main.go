package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ChunkedHandle 函数实现
func ChunkedHandle(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)

	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
			<-time.Tick(time.Second * 1)
		}
		close(ch)
	}()
	for {
		// 模拟消费者，有消息才会进入下一步
		v, ok := <-ch
		if !ok {
			flusher.Flush()
			return
		}
		fmt.Fprintf(w, "%v\n", v)
		flusher.Flush()
	}
}

func main() {
	http.HandleFunc("/chunked", ChunkedHandle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
