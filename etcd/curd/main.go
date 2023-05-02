package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	client "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"k8s.io/klog/v2"
)

type Client struct {
	client *client.Client
	err    error
}

func NewClientWithCert(caFile, certFile, keyFile string, endPoints []string) *Client {
	c := &Client{}

	pemCerts, err := ioutil.ReadFile(caFile)
	if err != nil {
		c.err = err
		return c
	}

	// etcd 启用了双向 TLS 认证，所以客户端证书同样需要加载
	etcdClientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		c.err = err
		return c
	}

	// 创建一个空的 CA Pool
	// 因为后续只会链接 Etcd 的 api 端点，所以此处选择使用空的 CA Pool，然后只加入 Etcd CA 既可
	// 如果期望链接其他 TLS 端点，那么最好使用 x509.SystemCertPool() 方法先 copy 一份系统根 CA
	// 然后再向这个 Pool 中添加自定义 CA
	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM(pemCerts)

	config := client.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Second * 5,
		// 自定义 CA 及 Client Cert 配置
		TLS: &tls.Config{
			RootCAs:      rootCertPool,
			Certificates: []tls.Certificate{etcdClientCert},
		},
	}

	etcdClient, err := client.New(config)
	if err != nil {
		c.err = err
		return c
	}
	return &Client{
		client: etcdClient,
		err:    nil,
	}
}

func (c *Client) GetWithContext(ctx context.Context, key string) (string, error) {
	response, err := c.client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	for _, kv := range response.Kvs {
		klog.Infof("create_revision: %v, mod_revision: %v, version: %v, \n", kv.CreateRevision, kv.ModRevision, string(kv.Value))
		return string(kv.Value), nil

	}
	return "", nil
}

func (c *Client) PutWithContext(ctx context.Context, key, value string) error {
	_, err := c.client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteWithContext(ctx context.Context, key string) error {
	//_, err := c.client.Delete(ctx, key, client.WithPrefix())
	_, err := c.client.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Watch(key string) (kvChan chan string) {
	// 监听 key 前缀的一组key的值
	kvChan = make(chan string)
	go func() {
		wChan := c.client.Watch(context.Background(), key, client.WithPrefix())
		for watchResp := range wChan {
			for _, event := range watchResp.Events {
				klog.Infof("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
				kvChan <- string(event.Kv.Value)
			}
		}
	}()
	return
}

// CreateLease 创建一个租约, 到期后key会被删除
func (c *Client) CreateLease(ctx context.Context, ttl int64, key, value string) error {
	leaseGrantResponse, err := c.client.Grant(ctx, ttl)
	if err != nil {
		return err
	}

	putResponse, err := c.client.Put(ctx, key, value, client.WithLease(leaseGrantResponse.ID))
	if err != nil {
		return err
	}
	klog.Infof("%+v\n", putResponse)
	return nil
}

func (c *Client) Lock(pfx string, fn func()) error {
	session, err := concurrency.NewSession(c.client)
	if err != nil {
		return err
	}
	defer session.Close()

	m := concurrency.NewMutex(session, pfx) // prefix
	if err := m.Lock(context.TODO()); err != nil {
		return err
	}

	fn()

	if err := m.Unlock(context.TODO()); err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
