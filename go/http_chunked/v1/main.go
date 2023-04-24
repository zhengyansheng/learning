package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ChunkedHandle(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "Hello %d\n", i)
		flusher.Flush()
		<-time.Tick(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/chunked", ChunkedHandle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
