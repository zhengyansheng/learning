# Ingress

## Install

```bash
# k apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

## ingress-nginx-controller
```bash
➜  ~ k -n ingress-nginx get pod        
NAME                                        READY   STATUS      RESTARTS      AGE
ingress-nginx-admission-create-b49w8        0/1     Completed   0             25h
ingress-nginx-admission-patch-bx46r         0/1     Completed   2             25h
ingress-nginx-controller-7545ff76f9-spfxg   1/1     Running     2 (18m ago)   25h

➜  ~ k exec -it ingress-nginx-controller-7545ff76f9-spfxg -- sh
/etc/nginx $ ps -ef
PID   USER     TIME  COMMAND
    1 www-data  0:00 /usr/bin/dumb-init -- /nginx-ingress-controller --election-id=ingress-nginx-leader --controller-class=k8s.io/ingress-nginx --ingress-class=nginx --configmap=ingress-nginx/ingress-nginx-controller --validating-webhook=:8443 --validating-webhook-cert
   10 www-data  0:00 /nginx-ingress-controller --election-id=ingress-nginx-leader --controller-class=k8s.io/ingress-nginx --ingress-class=nginx --configmap=ingress-nginx/ingress-nginx-controller --validating-webhook=:8443 --validating-webhook-certificate=/usr/local/cer
   25 www-data  0:00 nginx: master process /usr/bin/nginx -c /etc/nginx/nginx.conf
   30 www-data  0:00 nginx: worker process
   31 www-data  0:00 nginx: worker process
   32 www-data  0:00 nginx: worker process
   33 www-data  0:00 nginx: worker process
   34 www-data  0:00 nginx: cache manager process
  188 www-data  0:00 sh
  194 www-data  0:00 ps -ef

```


_**工作原理**_
1. _dumb-init_ 是一个简单的初始化进程，通常作为 Docker 容器中的第一个进程启动。它的主要作用是处理信号转发、清理子进程等，防止进程僵死（zombie processes）
2. _ingress-controller_ 控制器会动态监听 Kubernetes API中 Ingress，Service，Endpoints，ConfigMap等资源你的变化，并动态生成 Nginx 的配置文件 ;当检测到 Ingress 规则变更时，ingress-controller 控制器会重新生成 /etc/nginx/nginx.conf 配置文件，并 通过向 NGINX 主进程发送信号来触发配置文件的重新加载 nginx -s reload
