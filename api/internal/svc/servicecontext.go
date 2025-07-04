package svc

import (
	"bookstore/api/internal/config"
	"bookstore/api/internal/middleware"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Adder          adder.Adder
	Checker        checker.Checker
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Adder:          adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker:        checker.NewChecker(zrpc.MustNewClient(c.Check)),
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret, c.Auth.AccessExpire).Handle,
	}
}
