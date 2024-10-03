# envoy

```bash
docker pull envoyproxy/envoy:v1.22.2

docker run -rm -it -v $(pwd)/envoy-config-1.yaml:/envoy-custom.yaml -p 9901:9901 -p 10000:10000 envoyproxy/envoy:v1.22.2 -c /envoy-custom.yaml
curl localhost:10000


docker run -rm -it -v $(pwd)/envoy-config-2.yaml:/envoy-custom.yaml -p 9901:9901 -p 10000:10000 envoyproxy/envoy:v1.22.2 -c /envoy-custom.yaml
curl -H host:acem.com localhost:10000
curl -H host:acem.co localhost:10000

docker run -p 8080:80 nginxdemos/hello:plain-text
curl localhost:8080
```
