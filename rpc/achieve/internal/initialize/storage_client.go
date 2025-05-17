package initialize

import (
	storage "github.com/bearllflee/scholar-track/rpc/storage/storage_client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func ShouldNewFileClient() (storage.Storage, error) {
	storageClient, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "storage.rpc",
		},
	})
	if err != nil {
		return nil, err
	}
	sc := storage.NewStorage(storageClient)

	return sc, nil
}
