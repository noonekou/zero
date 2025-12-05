# Go-Zero 跨域配置说明

## 配置方式

> [!IMPORTANT]
> 在 go-zero v1.8+ 中,CORS 需要通过**代码方式**配置,而不是在 YAML 配置文件中配置。

### 代码配置(推荐)

在 `main.go` 文件中,使用 `rest.WithCors()` 选项:

```go
server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
```

## 已配置的服务

### 1. Admin API (端口 8888)
- 文件: `admin/admin.go`
- 已添加: `rest.WithCors("*")`

### 2. API 服务 (端口 8889)
- 文件: `api/api.go`
- 已添加: `rest.WithCors("*")`

## 配置示例

### 基本配置(允许所有域名)

```go
// 在 main.go 中
server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
```

### 指定单个域名

```go
server := rest.MustNewServer(c.RestConf, rest.WithCors("http://localhost:5173"))
```

### 指定多个域名

```go
server := rest.MustNewServer(c.RestConf, 
    rest.WithCors("http://localhost:5173", "https://example.com"))
```

## 生产环境建议

> [!WARNING]
> 在生产环境中,**不要**使用 `rest.WithCors("*")`,应该明确指定允许的域名。

### 推荐的生产配置

```go
// 生产环境 - 指定具体域名
server := rest.MustNewServer(c.RestConf, 
    rest.WithCors(
        "https://yourdomain.com",
        "https://www.yourdomain.com",
    ))
```

### 开发和生产环境分离

```go
// 根据环境变量决定 CORS 配置
var corsOrigins []string
if c.Mode == "dev" {
    corsOrigins = []string{"*"}
} else {
    corsOrigins = []string{
        "https://yourdomain.com",
        "https://www.yourdomain.com",
    }
}
server := rest.MustNewServer(c.RestConf, rest.WithCors(corsOrigins...))
```

## 注意事项

> [!IMPORTANT]
> 1. 当 `AllowCredentials: true` 时,`AllowOrigins` 不能设置为 `"*"`,必须指定具体域名
> 2. 修改配置后需要重启服务才能生效
> 3. 如果使用 Docker Compose,需要重新构建并启动容器

## 重启服务

### 使用 Docker Compose

```bash
# 重启 admin 服务
docker-compose -f docker-compose.admin.yml restart

# 重启 api 服务
docker-compose -f docker-compose.api.yml restart
```

### 直接运行

```bash
# 重启 admin 服务
cd admin
go run admin.go -f etc/admin-api.yaml

# 重启 api 服务
cd api
go run api.go -f etc/api-api.yaml
```

## 验证跨域配置

可以使用以下方式验证跨域配置是否生效:

### 1. 使用 curl 测试

```bash
curl -H "Origin: http://localhost:5173" \
     -H "Access-Control-Request-Method: POST" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS \
     http://localhost:8888/api/v1/your-endpoint \
     -v
```

### 2. 检查响应头

成功配置后,响应头应包含:
- `Access-Control-Allow-Origin`
- `Access-Control-Allow-Methods`
- `Access-Control-Allow-Headers`
- `Access-Control-Allow-Credentials`

## 常见问题

### Q: 配置后仍然报跨域错误?

**A:** 检查以下几点:
1. 确认服务已重启
2. 检查前端请求的域名是否在 `AllowOrigins` 列表中
3. 如果 `AllowCredentials: true`,确保 `AllowOrigins` 不是 `"*"`
4. 检查浏览器控制台的具体错误信息

### Q: 如何允许多个域名?

**A:** 在 `AllowOrigins` 中添加多个域名:
```yaml
AllowOrigins:
  - "http://localhost:5173"
  - "http://localhost:3000"
  - "https://example.com"
```

### Q: 开发环境和生产环境使用不同配置?

**A:** 使用不同的配置文件:
- 开发环境: `admin-api.dev.yaml`
- 生产环境: `admin-api.yaml`

启动时指定配置文件:
```bash
go run admin.go -f etc/admin-api.dev.yaml  # 开发环境
go run admin.go -f etc/admin-api.yaml      # 生产环境
```
