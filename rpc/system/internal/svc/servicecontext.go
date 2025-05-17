package svc

import (
	"github.com/bearllflee/scholar-track/rpc/system/internal/config"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/initialize"
	"github.com/bearllflee/scholar-track/rpc/system/internal/service"
)

type ServiceContext struct {
	Config                  config.Config
	CasbinService           *service.CasbinService
	ApiService              *service.ApiService
	RoleService             *service.RoleService
	MenuService             *service.MenuService
	DictionaryService       *service.DictionaryService
	DictionaryDetailService *service.DictionaryDetailService
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initialize.MustNewGorm(c.DataSource)
	global.DB = db
	return &ServiceContext{
		Config:                  c,
		CasbinService:           service.CasbinServiceApp,
		ApiService:              service.ApiServiceApp,
		RoleService:             service.RoleServiceApp,
		MenuService:             service.MenuServiceApp,
		DictionaryService:       service.DictionaryServiceApp,
		DictionaryDetailService: service.DictionaryDetailServiceApp,
	}
}
