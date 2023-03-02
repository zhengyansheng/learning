package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
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

func TestWatch(c *client.Client) {
	// 监听效果
	// 测试写入
	go func() {
		for {
			_, _ = c.Put(context.Background(), "/config/name", time.Now().String())
			time.Sleep(2 * time.Second)
		}
	}()
	wChan := c.Watch(context.Background(), "/config/name", client.WithPrefix()) // 监听key前缀的一组key的值
	for watchResp := range wChan {
		for _, event := range watchResp.Events {
			fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}

func main() {
	c, err := NewEtcdClient()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var key = "/xx/registry/namespaces/default/pods/monkey"

	// put kv 到 etcd
	_, err = c.Put(ctx, key, "apiVersion: v1\nkind: Pod ...")
	if err != nil {
		panic(err)
	}
	log.Println("put etcd success")

	// 从 etcd 查询
	response, err := c.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	for i, kv := range response.Kvs {
		log.Printf("i: %v, kv: %v\n", i, kv)
	}
	log.Println("get etcd success")

	// 从 etcd 删除 key
	_, err = c.Delete(ctx, key)
	if err != nil {
		panic(err)
	}

	// 从 etcd 删除前缀
	_, err = c.Delete(ctx, key, client.WithPrefix())
	if err != nil {
		return
	}
	log.Println("删除成功")

	// watch
	TestWatch(c)

}
