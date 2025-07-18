package svc

import (
	"bookstore/admin/internal/config"
	"bookstore/admin/internal/middleware"
	"bookstore/rpc/auth/client/adminauthservice"
	"bookstore/rpc/user/client/adminuserservice"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	Auth                 adminauthservice.AdminAuthService
	AdminUser            adminuserservice.AdminUserService
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:               c,
		Auth:                 adminauthservice.NewAdminAuthService(zrpc.MustNewClient(c.AuthConf)),
		AdminUser:            adminuserservice.NewAdminUserService(zrpc.MustNewClient(c.UserConf)),
		AuthMiddleware:       middleware.NewAuthMiddleware(c.Authorization.AccessSecret, c.Authorization.AccessExpire).Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware(conn).Handle,
	}
}
