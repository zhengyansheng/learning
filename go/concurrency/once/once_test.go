package once

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	r := NewResource()
	for i := 0; i < 10; i++ {
		conn := r.getConn()
		t.Logf("conn: %p", conn)
	}
	for i := 0; i < 10; i++ {
		conn2 := r.getConnByOnce()
		t.Logf("conn2: %p", conn2)
	}
}

type resource struct {
	lock sync.Mutex
	conn net.Conn
	once sync.Once
}

func NewResource() *resource {
	return &resource{
		lock: sync.Mutex{},
		once: sync.Once{},
	}
}

func (r *resource) getConn() net.Conn {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.conn != nil {
		return r.conn
	}

	conn, _ := net.DialTimeout("tcp", "www.baidu.com:80", time.Second*3)
	return conn
}

func (r *resource) getConnByOnce() net.Conn {
	r.once.Do(func() {
		conn, _ := net.DialTimeout("tcp", "www.baidu.com:80", time.Second*3)
		r.conn = conn
	})

	return r.conn
}
