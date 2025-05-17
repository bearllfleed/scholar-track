package main

import (
	"flag"
	"fmt"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/config"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/initialize"
	apiServer "github.com/bearllfleed/scholar-track/rpc/system/internal/server/apiservice"
	casbinServer "github.com/bearllfleed/scholar-track/rpc/system/internal/server/casbin"
	dictionaryServer "github.com/bearllfleed/scholar-track/rpc/system/internal/server/dictionaryservice"
	menuServer "github.com/bearllfleed/scholar-track/rpc/system/internal/server/menuservice"
	roleServer "github.com/bearllfleed/scholar-track/rpc/system/internal/server/role"
	userServer "github.com/bearllfleed/scholar-track/rpc/system/internal/server/user"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	initialize.Casbin()
	initialize.MustInitSnowflake(c.WorkerID)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		system.RegisterCasbinServer(grpcServer, casbinServer.NewCasbinServer(ctx))
		system.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		system.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		system.RegisterApiServiceServer(grpcServer, apiServer.NewApiServiceServer(ctx))
		system.RegisterMenuServiceServer(grpcServer, menuServer.NewMenuServiceServer(ctx))
		system.RegisterDictionaryServiceServer(grpcServer, dictionaryServer.NewDictionaryServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting system server at %s...\n", c.ListenOn)
	s.Start()
}
