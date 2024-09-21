package desigin

import (
	"testing"
	"time"
)

func TestFunctionOptionsMode(t *testing.T) {

	s := NewServer()
	t.Logf("%+v", s)

	s = NewServer(WithAddr("localhost"), WithPort(9000))
	t.Logf("%+v", s)

}

// 函数选项模式 适用于函数参数超过5个以上
type Server struct {
	Addr         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Timeout      time.Duration
}

type Option func(server *Server)

func NewServer(options ...Option) *Server {
	svr := &Server{
		Addr:         "127.0.0.1",
		Port:         8000,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		Timeout:      time.Second * 2,
	}

	for _, option := range options {
		option(svr)
	}

	return svr
}

func WithAddr(addr string) Option {
	return func(server *Server) {
		server.Addr = addr
	}
}

func WithPort(port int) Option {
	return func(server *Server) {
		server.Port = port
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.WriteTimeout = timeout
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}
