package main

import (
	"fmt"
	"time"
	
	capi "github.com/hashicorp/consul/api"
)

func main() {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		panic(err)
	}
	
	// 注册服务
	service := &capi.CatalogRegistration{
		//ID:      "11-22-12-1313-3131-9",
		Node:    "node2",
		Address: "1.1.1.1",
		Service: &capi.AgentService{
			Service: "redis4",
			Address: "2.2.2.2",
			Port:    6379,
		},
		Check: &capi.AgentCheck{
			//Node: "node1",
			//CheckID: "",
			Definition: capi.HealthCheckDefinition{
				HTTP:             "http://localhost:8080/health",
				IntervalDuration: time.Second * 5,
				TimeoutDuration:  time.Second * 5,
			},
		},
	}
	register, err := client.Catalog().Register(service, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("request time: %v\n", register.RequestTime)
	
	// 注册健康检查
	//opt := &capi.QueryOptions{}
	//catalogServices, _, err := client.Catalog().Service("1.1.1.1", "", opt)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(catalogServices)
	//for i, svc := range catalogServices {
	//	fmt.Printf("Service %d: %v\n", i, svc)
	//}
	//services, err := client.Agent().Services()
	//if err != nil {
	//	panic(err)
	//}
	//for i, svc := range services {
	//	fmt.Printf("Service %d: %v\n", i, svc)
	//}
	
	//client.Agent().CheckRegister(&capi.AgentCheckRegistration{
	//	//ID:   "redis-6309-2",
	//	Name: "my-service-check",
	//	//ServiceID: "my-service-id",
	//	AgentServiceCheck: capi.AgentServiceCheck{
	//		HTTP:     "http://localhost:8500/health",
	//		Interval: "10s",
	//	},
	//})
	
	//err = client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
	//	Name:    "redis",
	//	Address: "1.1.1.1",
	//	Port:    6379,
	//	Check: &capi.AgentServiceCheck{
	//		HTTP:     "http://localhost:8500/health",
	//		Interval: "10s",
	//	}})
	//if err != nil {
	//	panic(err)
	//}
}

//type ServiceRegistration struct {
//	ID      string   `json:"ID"`
//	Name    string   `json:"Name"`
//	Tags    []string `json:"Tags"`
//	Address string   `json:"Address"`
//	Port    int      `json:"Port"`
//	Node    string   `json:"Node"`
//}
//
//func main2() {
//	// 创建服务注册信息
//	reg := ServiceRegistration{
//		ID:      "unique-service-id",
//		Name:    "my-service",
//		Tags:    []string{"tag1", "tag2"},
//		Address: "127.0.0.1",
//		Port:    8080,
//		Node:    "my-node", // 指定节点名称
//	}
//
//	// 将服务注册信息转换为 JSON 格式
//	regJSON, err := json.Marshal(reg)
//	if err != nil {
//		log.Fatalf("Error marshaling registration data: %v", err)
//	}
//
//	// 向 Consul 的 Catalog API 发送注册服务的请求
//	resp, err := http.Put("http://consul-server-address:8500/v1/agent/service/register", bytes.NewBuffer(regJSON))
//	if err != nil {
//		log.Fatalf("Error registering service: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// 读取响应
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalf("Error reading response: %v", err)
//	}
//
//	fmt.Println("Service registered successfully:", string(body))
//}
