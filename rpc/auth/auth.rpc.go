package main

import (
	"flag"
	"fmt"

	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/config"
	adminauthserviceServer "bookstore/rpc/auth/internal/server/adminauthservice"
	apiauthserviceServer "bookstore/rpc/auth/internal/server/apiauthservice"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/auth.rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		auth.RegisterAdminAuthServiceServer(grpcServer, adminauthserviceServer.NewAdminAuthServiceServer(ctx))
		auth.RegisterApiAuthServiceServer(grpcServer, apiauthserviceServer.NewApiAuthServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
