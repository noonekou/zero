package main

import (
	"flag"
	"fmt"

	"bookstore/rpc/user/internal/config"
	adminuserserviceServer "bookstore/rpc/user/internal/server/adminuserservice"
	userserviceServer "bookstore/rpc/user/internal/server/userservice"
	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		user.RegisterAdminUserServiceServer(grpcServer, adminuserserviceServer.NewAdminUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
