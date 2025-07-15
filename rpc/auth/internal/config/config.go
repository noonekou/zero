package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DataSource    string
	Authorization struct {
		AccessSecret string
		AccessExpire int64
	}
}
