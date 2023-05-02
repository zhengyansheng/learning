package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

var (
	endPoints   []string = []string{"https://10.112.0.20:2379"}
	etcdCA               = "/Users/zhengyansheng/tls/etcd/ca.crt"
	etcdCert             = "/Users/zhengyansheng/tls/etcd/server.crt"
	etcdCertKey          = "/Users/zhengyansheng/tls/etcd/server.key"
)

func NewEtcdClient() (*client.Client, error) {
	// 为了保证 HTTPS 链接可信，需要预先加载目标证书签发机构的 CA 根证书
	etcdCA, err := ioutil.ReadFile(etcdCA)
	if err != nil {
		log.Fatal(err)
	}

	// etcd 启用了双向 TLS 认证，所以客户端证书同样需要加载
	etcdClientCert, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个空的 CA Pool
	// 因为后续只会链接 Etcd 的 api 端点，所以此处选择使用空的 CA Pool，然后只加入 Etcd CA 既可
	// 如果期望链接其他 TLS 端点，那么最好使用 x509.SystemCertPool() 方法先 copy 一份系统根 CA
	// 然后再向这个 Pool 中添加自定义 CA
	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM(etcdCA)

	config := client.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Second * 5,
		// 自定义 CA 及 Client Cert 配置
		TLS: &tls.Config{
			RootCAs:      rootCertPool,
			Certificates: []tls.Certificate{etcdClientCert},
		},
	}
	return client.New(config)
}

func getValue(ctx context.Context, c *client.Client, key string) string {
	getResponse, err := c.Get(ctx, key)
	if err != nil {
		return ""
	}
	for _, kv := range getResponse.Kvs {
		v, _ := strconv.Atoi(string(kv.Value))
		return fmt.Sprintf("%d", v)
	}
	return ""
}

func BatchUpdate(ctx context.Context, c *client.Client, key string) (string, error) {
	getResponse, err := c.Get(ctx, key)
	if err != nil {
		return "", err
	}
	var value int
	for _, kv := range getResponse.Kvs {
		v, _ := strconv.Atoi(string(kv.Value))
		value = v
		break
	}

	_, err = c.Put(context.Background(), key, fmt.Sprintf("%d", value+1))
	if err != nil {
		return "", err
	}

	fmt.Printf("before value: %d, after value: %v success\n", value, getValue(ctx, c, key))
	return "", nil
}

func main() {
	c, err := NewEtcdClient()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var key = "/xx/test/config"
	var w sync.WaitGroup
	for i := 0; i < 5; i++ {
		w.Add(1)
		go func(ctx context.Context, c *client.Client, key string) {
			defer w.Done()
			BatchUpdate(ctx, c, key)
		}(ctx, c, key)
	}

	w.Wait()
	fmt.Println("Done")

}
