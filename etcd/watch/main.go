package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

var (
	endPoints   []string = []string{"https://10.112.0.20:2379"}
	etcdCA               = "/Users/zhengyansheng/tls/etcd/ca.crt"
	etcdCert             = "/Users/zhengyansheng/tls/etcd/server.crt"
	etcdCertKey          = "/Users/zhengyansheng/tls/etcd/server.key"
)

func NewClient() (*client.Client, error) {
	// 为了保证 HTTPS 链接可信，需要预先加载目标证书签发机构的 CA 根证书
	etcdCA, err := ioutil.ReadFile(etcdCA)
	if err != nil {
		return nil, err
	}

	// etcd 启用了双向 TLS 认证，所以客户端证书同样需要加载
	etcdClientCert, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
	if err != nil {
		return nil, err
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

func main() {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}

	var k = "/config/name"

	// put kv
	go func() {
		for {
			now := time.Now().Format("2006-01-02 15:04:05")
			_, _ = c.Put(context.Background(), k, now)
			time.Sleep(time.Second * 3)
		}
	}()

	// watch kv
	watchChan := c.Watch(context.Background(), k, client.WithPrefix())
	for ch := range watchChan {
		for _, e := range ch.Events {
			fmt.Printf("type: %v, k: %q, v:%q\n", e.Type, e.Kv.Key, e.Kv.Value)
		}
	}

}
