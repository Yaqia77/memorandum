# docker创建容器

```sh
docker run --name memorandum -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=memorandum -e MYSQL_USER=username -e MYSQL_PASSWORD=root -p 3306:3306 -d mysql:8.0
```

# protoc生成go文件

```sh
protoc -I internal/service/pb internal/service/pb/*.proto --go_out=paths=source_relative:internal/service --go-grpc_out=paths=source_relative:internal/service

```