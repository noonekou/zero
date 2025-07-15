package svc

import (
	"bookstore/admin/internal/config"
	"bookstore/admin/internal/middleware"
	"bookstore/rpc/auth/authclient"
	"bookstore/rpc/user/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Auth           authclient.Auth
	User           userclient.User
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Auth:           authclient.NewAuth(zrpc.MustNewClient(c.AuthConf)),
		User:           userclient.NewUser(zrpc.MustNewClient(c.UserConf)),
		AuthMiddleware: middleware.NewAuthMiddleware(c.Authorization.AccessSecret, c.Authorization.AccessExpire).Handle,
	}
}
