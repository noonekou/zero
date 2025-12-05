#!/bin/bash

# 重新构建并启动 admin 服务
echo "重新构建 admin 服务..."
docker compose -f docker-compose.admin.yml up -d --build admin-service

echo "等待服务启动..."
sleep 3

echo "测试 CORS 配置..."
curl -H "Origin: http://localhost:5173" \
     -H "Access-Control-Request-Method: POST" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS \
     http://localhost:8888/v1/auth/logout \
     -v

echo ""
echo "如果看到 Access-Control-Allow-Origin 响应头,说明 CORS 配置成功!"
