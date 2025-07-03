package main

import (
	"flag"
	"fmt"

	"bookstore/api/internal/config"
	"bookstore/api/internal/handler"
	"bookstore/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf" // 导入 configcenter 包
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/logx"

	// 导入 rest 包
	// 导入 subscriber 包
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/bookstore-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if c.ConfigCenter.Etcd.Key != "" && len(c.ConfigCenter.Etcd.Hosts) > 0 {
		// 加载配置中心
		ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{Hosts: c.ConfigCenter.Etcd.Hosts, Key: c.ConfigCenter.Etcd.Key})
		v, er := ss.Value()
		if er != nil {
			logx.Errorf("GetConfig failed: %v", er)
		}
		if len(v) > 0 {
			cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
				Type: "yaml", // Configuration value type: json, yaml, toml
			}, ss)

			var err error
			c, err = cc.GetConfig()
			if err != nil {
				// panic(err)
				logx.Errorf("GetConfig failed: %v", err)
			}

			cc.AddListener(func() {
				c, err = cc.GetConfig()
				if err != nil {
					logx.Errorf("GetConfig failed -+: %v", err)
					// panic(err)
				}
			})
		}
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
