# 网络

## 

## 1. OSI

> OSI 7层模型
>
> 分层的目的是解偶

在互联网上通信时，IP在数据包中是不变的，代表的是端点。

但是在通信的过程中，数据包的MAC地址一直被替换，MAC地址代表的是下一跳节点。



### 



## 2. HTTP

### 1. keepalive



## 3. TCP

> TCP 是一种*面向连接的*、*可靠的*、*基于字节流的*传输层通信协议。

- 面向连接的：三次握手完成后，到开辟完资源，这个连接就有了。（通过彼此心跳 PING/PONG 来维持彼此的存活）
- 可靠的：每一次发送报文都要得到对方的回复确认，来保证消息是可靠的，如果没有回复确认，发送方会一直发送 请求报文。
- 基于字节流的：



TCP在传输时是有序的

TCP是传输控制，控制两端建立可靠的连接，包的发送频率等，并不负责包的发送。
TCP本质上就是建立在socket以上的，socket有ip和port组成。



### 1. TCP 状态机

<img src="imgs/image-20210921100646389.png" alt="image-20210921100646389" style="zoom: 50%;" />





#### 1. 常见的信号

1) SYN: 请求建立连接

2. FIN:   断开连接的请求

3. ACK: 请求已确认
4. RST: 非正常的链接，未经过四次挥手





### 2. 三次握手

#### 1. tcpdump 抓包

```bash
# tcpdump -i eth0 port 8000 -nn -c 3 -w tcp_demo.pcap

# tcpdump -nn -i eth0 port 8000 or arp -w tcp_demo.pcap
```



```
参数
-i     监听的网卡
port 监听的端口
-c     抓取三个包后退出
-w    写入到文件中
-n     不解析端口号为协议名
```

#### 2. Wireshark 分析报文

![image-20210920185254711](imgs/image-20210920185254711.png)





### 3. 四次挥手 或 三次挥手

> 为了释放，回收资源



从TCP协议来看，挥手必须要四次，但是第三次和第四次是可以合并的。





<img src="imgs/image-20210921100215431.png" alt="image-20210921100215431" style="zoom: 33%;" />







<img src="imgs/image-20210921110030219.png" alt="image-20210921110030219" style="zoom:50%;" />



由于TCP是全双工的，前2次挥手是Client -> Server的半关闭，后2次挥手是Server -> Client的半关闭，合起来才是真正的关闭连接。

这里的半关闭并不是真正意义上的关闭

主动关闭的一方会进入 TIME-WAIT 状态







#### 1. TIME-WAIT -> CLOSED

**MSL (Maximum Segment Lifetime)**

一个TCP报文在网络上存活的最大时间



**维持2MSL时长的 TIME-WAIT 状态**

- 保证至少一次报文的往返时间内端口是不可用的



要等2MSL(Maximum Segment Life)时间，MSL是一个TCP报文在网络上存活的最长时间。

主动关闭的一方会主动进入TIME-WAIT状态，通常会保持2分钟左右，端口是被占用的。



#### 2. TIME-WAIT 优化

1) net.ipv4.tcp_tw_reuse = 1

开启后，作为客户端时新连接可以使用仍然处于 TIME-WAIT 状态的端口。

由于 timestamp 的存在，操作系统可以拒绝迟到的报文。 net.ipv4.tcp_timestamps = 1 这个也同时开启



2. net.ipv4.tcp_tw_recycle = 0

开启后，同时作为客户端 和 服务器 都可以使用 TIME-WAIT 状态的端口

不安全，无法避免报文延迟，重复等给新连接造成混乱



3. net.ipv4.tcp_max_tw_buckets = 262144

time_wait 状态连接的最大数量

超出后直接关闭连接



#### 3. Wireshark 分析三次挥手

![image-20210921111012472](imgs/image-20210921111012472.png)



从截图来看，确实是三次挥手，而不是四次挥手。





### 



### 4. 滑动窗口

TCP 是按Seq序列号接收数据的



接收窗口：发送方维护接收窗口，表示接收方可用的缓存空间。











### 5. 拥塞控制





![image-20210925215338248](imgs/image-20210925215338248.png)



#### 1. 慢启动

<img src="imgs/image-20210925180036036.png" alt="image-20210925180036036" style="zoom:50%;" />

#### 2. 拥塞避免

<img src="imgs/image-20210925180701869.png" alt="image-20210925180701869" style="zoom: 33%;" />



#### <img src="imgs/image-20210925175943863.png" alt="image-20210925175943863" style="zoom: 33%;" /> 



#### 3. 快速重传



#### 4. 快速恢复



### 6. 粘包



### x. 常见问题

#### 1. SYN中Seq为什么不同

>  三次握手中，为什么要Seq到序列号都要+1，而不是用相同的一个序列号？

因为报文在网络上传输会有延迟，会有丢失，会有重传机制，为了区分这些问题所以序列号要一直累加不能重复。



#### 2. mtu 是什么



#### 3. 为什么握手要三次而不是二次

因为TCP是全双工的，任意一方都可以发送和接收数据。

站在客户端一方来看，客户端发送SYN，并得到服务端回复了ACK，此时表示客户端的输入和输出都是通的。

站在服务端一方来看，服务端回复了客户端的ACK，但是自已发送的SYN并没有得到回复，此时仅能表示服务端输入是通的，输出无法保证。

因此握手必须三次，才能保证彼此的输入和输出是通的。



#### 4. 为什么挥手是四次而不是三次

假设： Client -> Server

因为TCP是可靠的，Client发送FIN报文，必须收到Server回复了ACK，才算可靠的，否则Client就要一直发送FIN报文，直到收到Server回复了ACK



**如果四次挥手，改成三次，把中间的第二，三次合并成一次，是否可行**

首先第二次挥手是Server回复ACK，第三次挥手是Server发送FIN报文



- 先来解释 为什么要有第二次挥手和第三次发送FIN

第二次挥手，是因为Client发送了FIN请求，要断开连接，Server必须回复ACK，但是此时Server端可能还有未处理的数据，等Server端数据处理完成，发送FIN请求断开连接，此处处于LAST_ACK状态，等待Client回复ACK，此时Server处于CLOSE状态。



如果把第二次的ACK和第三次的FIN合并，会导致Client接收到Server回复ACK延迟(因为Server要处理断开后要处理的数据)，Client没有收到ACK，会再次发送FIN

## 4. IP

> 硬件设备: 路由器
>
> 网络层就是寻找下一跳





## 5. MAC

> 硬件设备: 交换机

```
# apr -a
```



## 6. 工具

### 1. tcpdump



### 2. Wireshark

<img src="imgs/image-20210925163702834.png" alt="image-20210925163702834" style="zoom:50%;" />

## 


