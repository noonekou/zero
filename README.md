# zero

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
