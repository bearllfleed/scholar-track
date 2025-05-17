package initialize

import "github.com/bearllfleed/scholar-track/rpc/system/internal/service"

var casbinService = service.CasbinServiceApp

func Casbin() {
	casbinService.Casbin()
}
