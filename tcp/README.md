# TCP
> Transfer Controller Protocol  
> 面向连接的，可靠的网络传输协议。

- 顺序控制
- 重发控制
- 流量控制
- 拥塞控制

## 五元组

```text

c [源IP:源端口:目标IP:目标端口:协议] --> S


```


## 三次握手 四次回首


## 重发控制

重发的时间如何衡量  
- 初始发送时，重发的超时时间设置为6秒
- 非初始发送时，发送数据，回应数据，这一来一回的时间可以作为重发的超时时间


## 流量控制

TCP首部有Windows Size字段来标识服务端可接收的缓冲区大小。  
服务端会自动的控制这个Windows Size字段的大小自动调整做到流量控制。

## 拥塞控制

慢启动   
发送1MSS  
发送2MSS  
发送3MSS  
发送4MSS  
超时，未收到ACK  
慢启动  
发送1MSS  
发送2MSS  
发送3MSS  
发送4MSS  
发送5MSS  
超过阈值  
慢启动  
发送1MSS  
发送2MSS   
发送3MSS  
发送4MSS  
发送5MSS  




# UDP
> User Datagram Protocol
> 面向无连接的，不具备可靠传输的数据报协议。


## 使用场景

- DHCP
- DNS
- RIP
---
- IP电话
- 视频电话