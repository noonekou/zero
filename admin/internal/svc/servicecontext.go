package svc

import (
	"bookstore/admin/internal/config"
	"bookstore/admin/internal/middleware"
	"bookstore/rpc/auth/client/adminauthservice"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Auth           adminauthservice.AdminAuthService
	AdminUser      adminuserservice.AdminUserService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Auth:           adminauthservice.NewAdminAuthService(zrpc.MustNewClient(c.AuthConf)),
		AdminUser:      adminuserservice.NewAdminUserService(zrpc.MustNewClient(c.UserConf)),
		AuthMiddleware: middleware.NewAuthMiddleware(c.Authorization.AccessSecret, c.Authorization.AccessExpire).Handle,
	}
}
