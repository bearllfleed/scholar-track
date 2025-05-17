package svc

import (
	"github.com/bearllflee/scholar-track/rpc/storage/internal/config"
	"github.com/bearllflee/scholar-track/rpc/storage/internal/initialize"
	"github.com/bearllflee/scholar-track/rpc/storage/internal/service"
)

type ServiceContext struct {
	Config                 config.Config
	StorageService         *service.StorageService
	FileService            *service.FileService
	FailedDeletionsService *service.FailedDeletionsService
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initialize.MustNewGrom(c.DataSource)
	storageService, err := service.NewStorageService(c.Storage)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                 c,
		StorageService:         storageService,
		FileService:            service.NewFileService(db),
		FailedDeletionsService: service.NewFailedDeletionsService(db),
	}
}
