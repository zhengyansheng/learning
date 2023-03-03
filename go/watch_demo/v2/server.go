package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	addr     = ":8000"
	notifyCh = make(chan int)
)

// Production 生产者
func Production() {
	for i := 0; i < 5; i++ {
		notifyCh <- i + 1
		<-time.Tick(time.Second * 3)
	}
	close(notifyCh)
}

// Watch 函数实现
func Watch(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)
	for {
		// 模拟消费者，有消息才会进入下一步
		v, ok := <-notifyCh
		if !ok {
			flusher.Flush()
			return
		}
		fmt.Fprintf(w, "%v\n", v)
		flusher.Flush()
	}
}

func main() {
	go Production()
	http.HandleFunc("/watch", Watch)
	fmt.Printf("listen addr: %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
