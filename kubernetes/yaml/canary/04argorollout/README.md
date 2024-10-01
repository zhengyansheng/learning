# 灰度发布

## 蓝绿发布


## 金丝雀发布  
```bash
# k get rollouts

# k argo rollouts promote <rollout-canary>
```



## rollout canary
```bash

# vim /etc/hosts
127.0.0.2 shark.local


# while true
do
curl http://shark.local
sleep 1
done


# curl -H "X-Canary: always" shark.local
```
