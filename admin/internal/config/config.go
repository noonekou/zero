package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Authorization struct {
		AccessSecret string
		AccessExpire int64
	}
	AuthConf   zrpc.RpcClientConf
	UserConf   zrpc.RpcClientConf
	DataSource string
	TokenRedis redis.RedisConf
}
