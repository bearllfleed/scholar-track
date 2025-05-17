package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bearllflee/scholar-track/rpc/storage/internal/config"
	"github.com/bearllflee/scholar-track/rpc/storage/internal/server"
	"github.com/bearllflee/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/storage/internal/task"
	"github.com/bearllflee/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/storage.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		storage.RegisterStorageServer(grpcServer, server.NewStorageServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	failedDeletionTask := task.NewClearFailedDeletionsTask(ctx)
	go failedDeletionTask.Start(context.Background())
	s.Start()
}
