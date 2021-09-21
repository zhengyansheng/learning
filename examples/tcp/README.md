# 网络

## 

## 1. TCP

### 1. TCP状态机

![image-20210921100646389](imgs/image-20210921100646389.png)







### 2. 三次握手

#### 1. tcpdump 抓包

```bash
# tcpdump -i eth0 port 8000 -c 3 -w tcp_demo.cap
```



```
参数
-i     监听的网卡
port 监听的端口
-c     抓取三个包后退出
-w    写入到文件中
```

#### 2. Wireshark 分析报文

![image-20210920185254711](imgs/image-20210920185254711.png)





### 2. 四次挥手 或 三次挥手

![image-20210921100215431](imgs/image-20210921100215431.png)







![image-20210921110030219](imgs/image-20210921110030219.png)



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





### 3. 拥塞控制



### 4. 滑动窗口

TCP 是按Seq序列号接收数据的







### x. 常见问题

#### 1. SYN中Seq为什么不同

>  三次握手中，为什么要Seq到序列号都要+1，而不是用相同的一个序列号？

因为报文在网络上传输会有延迟，会有丢失，会有重传机制，为了区分这些问题所以序列号要一直累加不能重复。



#### 2. mtu 是什么





## 2. HTTP

