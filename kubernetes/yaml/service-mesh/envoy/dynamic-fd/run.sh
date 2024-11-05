docker run --rm -it -v $(pwd)/envoy-config.yaml:/envoy-custom.yaml \
  -v $(pwd)/lds.yaml:/var/lib/envoy/lds.yaml -v $(pwd)/cds.yaml:/var/lib/envoy/cds.yaml \
  -p 9901:9901 -p 10000:10000 -p 19000:19000 envoyproxy/envoy:v1.22.2 -c /envoy-custom.yaml --log-level debug
