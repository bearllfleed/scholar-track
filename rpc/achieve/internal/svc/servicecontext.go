package svc

import (
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/config"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/initialize"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/service"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage_client"
)

type ServiceContext struct {
	Config          config.Config
	CategoryService *service.CategoryService
	PropertyService *service.PropertyService
	AchieveService  *service.AchieveService
	StorageClient   storage_client.Storage
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initialize.MustNewGorm(c.DataSource)
	return &ServiceContext{
		Config:          c,
		CategoryService: service.NewCategoryService(db),
		PropertyService: service.NewPropertyService(db),
		AchieveService:  service.NewAchieveService(db),
	}
}
