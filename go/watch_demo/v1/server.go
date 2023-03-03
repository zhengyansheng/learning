package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Watch(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)
	for i := 0; i < 5; i++ {
		fmt.Fprintf(w, "Hello World\n")
		flusher.Flush()
		<-time.Tick(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/watch", Watch)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
