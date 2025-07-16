package svc

import (
	"bookstore/api/internal/config"
	"bookstore/api/internal/middleware"
	"bookstore/rpc/auth/client/apiauthservice"
	"bookstore/rpc/user/client/userservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Auth           apiauthservice.ApiAuthService
	User           userservice.UserService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c.Authorization.AccessSecret, c.Authorization.AccessExpire).Handle,
		Auth:           apiauthservice.NewApiAuthService(zrpc.MustNewClient(c.AuthConf)),
		User:           userservice.NewUserService(zrpc.MustNewClient(c.UserConf)),
	}
}
