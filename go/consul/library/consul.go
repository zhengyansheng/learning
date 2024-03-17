package library

import (
	"fmt"
	"time"
	
	"github.com/google/uuid"
	capi "github.com/hashicorp/consul/api"
)

// consulClient 定义consul client
type consulClient struct {
	client  *capi.Client
	ClientX *capi.Client
}

// NewConsulClient 初始化 consul client
func NewConsulClient() *consulClient {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &consulClient{client: client, ClientX: client}
}

// RegisterService 注册服务和健康检查
func (c *consulClient) RegisterService(reg *capi.CatalogRegistration, q *capi.WriteOptions) error {
	_, err := c.client.Catalog().Register(reg, q)
	return err
}

func (c *consulClient) DeregisterService(dc, node, serviceID, checkID string, q *capi.WriteOptions) error {
	deReg := &capi.CatalogDeregistration{
		Datacenter: dc,
		Node:       node,
		ServiceID:  serviceID,
		CheckID:    checkID,
	}
	_, err := c.client.Catalog().Deregister(deReg, q)
	return err
}

// GeneratorCatalogRegistration 生成注册服务的内容
func (c *consulClient) GeneratorCatalogRegistration(dc, node, address string, service ServiceOption, check CheckOption) *capi.CatalogRegistration {
	// https://developer.hashicorp.com/consul/api-docs/catalog
	uid, _ := uuid.NewUUID()
	
	agentService := &capi.AgentService{
		ID:      service.Name,
		Service: service.Name,
		Address: service.Address,
		Port:    service.Port,
		Tags:    service.Tags,
	}
	
	tcpHealthCheck := capi.HealthCheckDefinition{
		TCPUseTLS: false,
		TCP:       "www.baidu.com:80",
		Interval:  *capi.NewReadableDuration(time.Second * 3),
		Timeout:   *capi.NewReadableDuration(time.Second * 2),
	}
	
	return &capi.CatalogRegistration{
		Datacenter: dc,
		ID:         uid.String(),
		Node:       node,
		Address:    address,
		Service:    agentService,
		Check: &capi.AgentCheck{
			Node:       node,
			ServiceID:  service.Name,
			CheckID:    service.Name,
			Name:       fmt.Sprintf("%v Health Status", service.Name),
			Definition: tcpHealthCheck,
		},
	}
}

func (c *consulClient) AgentRegisterService(name string) error {
	service := &capi.AgentServiceRegistration{
		Name:    name,
		Address: "127.0.0.1",
		Port:    8500,
		Tags:    []string{"1", "2"},
		Check: &capi.AgentServiceCheck{
			TCP:      "127.0.0.1:8500",
			Interval: "3s", // 定期检查的时间间隔
			Timeout:  "2s", //
		},
	}
	return c.client.Agent().ServiceRegister(service)
}
