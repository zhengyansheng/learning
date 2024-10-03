# Istio

## 安装
```bash
# wget https://github.com/istio/istio/releases/download/1.23.2/istioctl-1.23.2-osx-arm64.tar.gz 

# istioctl install --set profile=demo -y
        |\          
        | \         
        |  \        
        |   \       
      /||    \      
     / ||     \     
    /  ||      \    
   /   ||       \   
  /    ||        \  
 /     ||         \ 
/______||__________\
____________________
  \__       _____/  
     \_____/        

✔ Istio core installed ⛵️                                                                                                                                                                                                                                                      
✔ Istiod installed 🧠                                                                                                                                                                                                                                                         
✔ Egress gateways installed 🛫                                                                                                                                                                                                                                                
✔ Ingress gateways installed 🛬                                                                                                                                                                                                                                               
✔ Installation complete                                                                                                                                                                                                                                                       Made this installation the default for cluster-wide operations.


```

## 自动注入
```bash
# k create ns ops

# k label namespace ops istio-injection=enabled
```

## pilot-discovery
```bash
➜  ~ k -n istio-system get pods
NAME                                    READY   STATUS    RESTARTS   AGE
istio-egressgateway-f89d474b8-6fs4k     1/1     Running   0          3m4s
istio-ingressgateway-7f458c4596-kgkb7   1/1     Running   0          3m4s
istiod-6ff6fc8cdb-cl4z5                 1/1     Running   0          3m23s

➜  ~ k -n istio-system exec -it istiod-6ff6fc8cdb-cl4z5 -- ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
istio-p+       1       0  0 04:10 ?        00:00:00 /usr/local/bin/pilot-discovery discovery --monitoringAddr=:15014 --log_output_level=default:info --domain cluster.local --keepaliveMaxServerConnectionAge 30m
istio-p+      18       0  0 04:13 pts/0    00:00:00 ps -ef

```

_pilot-discovery_  
是 Istio 中的控制平面组件之一，负责管理服务发现和配置。  
它监控 Kubernetes 中的service和endpoint，确保 Envoy 代理能够获取所需的配置信息，从而实现流量管理、路由和策略应用。  


