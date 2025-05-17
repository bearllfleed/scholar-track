package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/config"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/initialize"
	achieveservice "github.com/bearllfleed/scholar-track/rpc/achieve/internal/server/achieveservice"
	categoryservice "github.com/bearllfleed/scholar-track/rpc/achieve/internal/server/categoryservice"
	propertyservice "github.com/bearllfleed/scholar-track/rpc/achieve/internal/server/propertyservice"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/achieve.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	initialize.MustInitSnowflake(c.WorkerID)

	sc, err := initialize.ShouldNewFileClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx.StorageClient = sc
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		achieve.RegisterCategoryServiceServer(grpcServer, categoryservice.NewCategoryServiceServer(ctx))
		achieve.RegisterPropertyServiceServer(grpcServer, propertyservice.NewPropertyServiceServer(ctx))
		achieve.RegisterAchieveServiceServer(grpcServer, achieveservice.NewAchieveServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
