package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add          zrpc.RpcClientConf
	Check        zrpc.RpcClientConf
	ConfigCenter struct {
		Etcd struct {
			Hosts []string
			Key   string
		}
	}
}
