package desigin

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func TestDecoratorMode(t *testing.T) {

	// 针对路由函数装饰器
	Run()

	// 针对mux装饰
	Run2()
}

// 装饰器模式1
type Handler func(http.ResponseWriter, *http.Request)

func DecoratorHandle(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler(w, r)
		log.Printf("url: %v, duration: %v", r.URL, time.Since(now))
	}
}

// 装饰器模式2
func HandleMiddle(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("url: %v, duration: %v", r.URL, time.Since(now))
	}
	return http.HandlerFunc(fn)
}

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /say", DecoratorHandle(Say))
	mux.HandleFunc("GET /hello", Hello)

	s := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}
	err := s.ListenAndServe()
	log.Fatal(err)
}

func Say(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("say"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

func Run2() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /say", DecoratorHandle(Say))
	mux.HandleFunc("GET /hello", Hello)

	s := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: HandleMiddle(mux),
	}
	err := s.ListenAndServe()
	log.Fatal(err)
}
