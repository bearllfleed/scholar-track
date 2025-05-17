package initialize

import "github.com/bearllflee/scholar-track/rpc/system/internal/service"

var casbinService = service.CasbinServiceApp

func Casbin() {
	casbinService.Casbin()
}
