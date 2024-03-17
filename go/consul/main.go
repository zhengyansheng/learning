package main

import (
	"github.com/zhengyansheng/consul/library"
)

func main() {
	client := library.NewConsulClient()
	
	// 注册服务
	//service := client.GeneratorCatalogRegistration(
	//	"dc1", "consul-agent-01", "172.17.0.1",
	//	library.ServiceOption{Name: "ssh3", Address: "172.17.0.1", Port: 8500, Tags: []string{"dev", "v1"}},
	//	library.CheckOption{},
	//)
	//err := client.RegisterService(service, nil)
	//if err != nil {
	//	panic(err)
	//}
	//klog.Infof("register service %v ok", service.Service.Service)
	//
	// register
	err := client.AgentRegisterService("php4")
	if err != nil {
		panic(err)
	}
}
