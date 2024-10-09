# docker-compose

## 命令
```bash
# docker-compose up

# docker-compose down
```

## 检测
```bash
mysqladmin ping -uroot -proot123  -h 127.0.0.1

curl -X GET "localhost:9200/_cluster/health?pretty"

docker exec -it redis redis-cli ping

d docker exec -it influxdb influx

x.x.x.x:18083
```
