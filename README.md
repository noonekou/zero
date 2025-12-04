# 项目搭建

## 步骤

### 修改 xxx.api
执行 `goctl api go -api xxx.api -dir .` 同步

### 修改rpc xxx.proto
执行 `goctl rpc protoc xxx.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m` 同步

### 数据同步 
执行 `goctl model pg datasource -url="postgres://postgres:123456@localhost:5432/gozero?sslmode=disable" -table=t_role_permission -dir=.` 同步

### 数据库连接配置, 修改api权限sql
`psql -U postgres -d gozero`

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

curl -X POST http://localhost:8888/v1//auth/login -H 'Content-Type: application/json' -d '{"username":"admin","password":"21223e1706c109dca4af2c7b1f2fff69"}'

curl -X POST http://localhost:8888/v1//auth/login -H 'Content-Type: application/json' -d '{"username":"admin@example.com","password":"21223e1706c109dca4af2c7b1f2fff69"}'


curl -i "http://localhost:8888/v1/user/info"

# 测试rpc服务
# yaml 配置Mode = dev
grpcui -plaintext 127.0.0.1:8080
# 或者
grpcurl -plaintext 127.0.0.1:8080 user.User/GetUserInfo

# docker 访问rpc
grpcurl -plaintext -proto rpc/auth/auth.rpc.proto localhost:8180

# pgsql 生成model
goctl model pg datasource -url="postgres://localhost:5432/test?sslmode=disable" -table=t_admin_user,t_admin_user_role,t_api_permission,t_apis,t_permission,t_resource,t_role,t_role_permission,t_user -dir=.

goctl model pg datasource -url="postgres://localhost:5432/test?sslmode=disable" -table=t_role -dir=.

# 容器内执行 (host: pg)
goctl model pg datasource -url="postgres://postgres:123456@localhost:5432/gozero?sslmode=disable" -table=t_role -dir=.

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
# 分组 需要加 -m
goctl rpc protoc xxx.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m
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
# 指定配置文件
docker-compose -f docker-compose.dev.yml up -d --build

# 指定服务
docker-compose -f docker-compose.admin.yml up -d --build admin-service
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

6. 进入容器 pg

```Dockerfile
docker compose -f docker-compose.admin.yml exec pg psql -U postgres gozero


docker compose exec [服务名称] bash
```

# 监控

1. prometheus

http://localhost:9095/

2. jaeger

http://localhost:16686/

# ETCD Keeper

http://localhost:8999/
http://localhost:8999/etcdkeeper/?endpoint=http://etcd:2379


```bash
docker pull evildecay/etcdkeeper
docker run -d -p 8999:8080 --name etcdkeeper evildecay/etcdkeeper
# 或者 指定 etcd 地址
docker run -d -p 8999:8000 --name etcdkeeper \
  -e ETCD_URL=http://host.docker.internal:2379 \
  evildecay/etcdkeeper
# 或者 指定 容器地址
docker run -d -p 8999:8000 --name etcdkeeper \
  -e ETCD_URL=http://host.docker.internal:2379 \
  evildecay/etcdkeeper
```

# PG

```pg
// 连接本地数据库
psql -U username -d dbname

// 连接远程数据库
psql -h hostname -p port -U username -d dbname
```

常用命令
\l 列出所有数据库
\c 连接数据库
\dt 列出当前数据库中的所有表
\d 表的结构
\q 退出

# Redis

```redis
// 连接本地数据库
redis-cli -a 123456 -h 127.0.0.1 -p 6379

// 连接远程数据库
redis-cli -h hostname -p port -a 123456 -h 127.0.0.1 -p 6379

docker exec -it redis redis-cli -a 123456
```

常用命令
KEYS * 列出所有key
DEL key 删除key
GET key 获取key的值
SET key value 设置key的值
EXISTS key 检查key是否存在
EXPIRE key seconds 设置key的过期时间
TTL key 查看key的剩余生存时间
EXPIREAT key timestamp 设置key的过期时间

INFO 查看 Redis 服务器的详细信息（如内存、客户端连接数等）
FLUSHDB 清空数据库
FLUSHALL 清空所有数据库

# API 文档

```bash
sh doc-sync.sh
```
