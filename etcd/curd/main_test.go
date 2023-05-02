package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"k8s.io/klog/v2"
)

var (
	endPoints   []string = []string{"https://10.112.0.20:2379"}
	etcdCA               = "/Users/zhengyansheng/tls/etcd/ca.crt"
	etcdCert             = "/Users/zhengyansheng/tls/etcd/server.crt"
	etcdCertKey          = "/Users/zhengyansheng/tls/etcd/server.key"
)

func TestGet(t *testing.T) {
	c := NewClientWithCert(etcdCA, etcdCert, etcdCertKey, endPoints)
	if c.err != nil {
		t.Fatal(c.err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key := "/xx/registry/config"
	getWithContext, err := c.GetWithContext(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	klog.Infof("key: %v, value: %v", key, getWithContext)
}

func TestPut(t *testing.T) {
	c := NewClientWithCert(etcdCA, etcdCert, etcdCertKey, endPoints)
	if c.err != nil {
		t.Fatal(c.err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key := "/xx/registry"
	err := c.PutWithContext(ctx, key, "1234")
	if err != nil {
		t.Fatal(err)
	}
	klog.Infof("key: %v", key)
}

func TestWatch(t *testing.T) {
	c := NewClientWithCert(etcdCA, etcdCert, etcdCertKey, endPoints)
	if c.err != nil {
		t.Fatal(c.err)
	}

	key := "/xx/registry/config"
	for value := range c.Watch(key) {
		klog.Info(value)
	}
}

func TestLease(t *testing.T) {
	c := NewClientWithCert(etcdCA, etcdCert, etcdCertKey, endPoints)
	if c.err != nil {
		t.Fatal(c.err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()

	key := "/xx/registry/config"
	err := c.CreateLease(ctx, 15, key, "12345")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for {
			getWithContext, err := c.GetWithContext(ctx, key)
			if err != nil {
				t.Fatal(err)
			}
			klog.Infof("key: %v, value: %v", key, getWithContext)
			time.Sleep(time.Second * 1)
		}
	}()
	select {}
}

func TestLock(t *testing.T) {
	c := NewClientWithCert(etcdCA, etcdCert, etcdCertKey, endPoints)
	if c.err != nil {
		t.Fatal(c.err)
	}

	w := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		w.Add(1)
		go func(w *sync.WaitGroup, i int) {
			defer w.Done()
			c.Lock("/my-test-lock", func() {
				fmt.Println("start", i)
				time.Sleep(time.Second * 3)
				fmt.Println("end", i)
			})
		}(w, i)
	}
	w.Wait()
	fmt.Println("done")
}
