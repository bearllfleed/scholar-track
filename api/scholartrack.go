package main

import (
	"flag"
	"fmt"
	"github.com/bearllflee/scholar-track/api/internal/config"
	"github.com/bearllflee/scholar-track/api/internal/handler"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/scholartrack-api.yaml", "the config file")

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &config.GlobalConfig)

	server := rest.MustNewServer(config.GlobalConfig.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(config.GlobalConfig)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", config.GlobalConfig.Host, config.GlobalConfig.Port)
	server.Start()
}
