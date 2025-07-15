# 项目搭建

## 创建 admin 服务

```bash
goctl api new xxx
```

## 根据 xxx.api 生成 xxx-api 服务

```bash
goctl api go -api xxx.api -dir .

# 运行 xxx-api 服务
go run xxx.go -f etc/xxx-api.yaml

# 测试 xxx-api 服务
curl -X POST http://localhost:8888/v1/user/login -H 'Content-Type: application/json' -d '{"username":"admin","password":"123456"}'

curl -i "http://localhost:8888/v1/user/info"

# 测试rpc服务
# yaml 配置Mode = dev
grpcui -plaintext 127.0.0.1:8080
# 或者
grpcurl -plaintext 127.0.0.1:8080 user.User/GetUserInfo

# pgsql 生成model
goctl model pg datasource -url="postgres://localhost:5432/test?sslmode=disable" -table=t_user,t_role,t_admin_user,t_permission,t_role_permission,t_user_role -dir=.

# pg 连接配置 etc/xxx.yaml
# config
# svr context
```

## 创建 rpc 服务

```bash
# 使用命令生成proto文件模版
goctl rpc -o xxx.proto

goctl rpc new xxx

# 根据 xxx.proto 生成 xxx-rpc 服务
goctl rpc protoc xxx.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## 修改 api 配置文件，实现 rpc 调用

```yaml
XXX:
  Etcd:
    Hosts:
      - localhost:2379
    Key: xxx.rpc
```

# run and deploy

1. 构建并启动所有服务 (首次运行或代码/配置有变化时)：

```Dockerfile
docker-compose up -d --build
```

2. 启动所有服务：

```Dockerfile
docker-compose up -d
```

3. 查看服务状态

```Dockerfile
docker-compose ps
```

4. 查看服务日志

```Dockerfile
docker-compose logs -f api-service
docker-compose logs -f add-rpc-service
docker-compose logs -f check-rpc-service
```

5. 停止所有服务：

```Dockerfile
docker-compose down
```

# 监控

1. prometheus

http://localhost:9090/

2. jaeger

http://localhost:16686/
