package svc

import (
	"github.com/bearllflee/scholar-track/rpc/achieve/client/achieveservice"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/config"
	"github.com/bearllflee/scholar-track/api/internal/middleware"
	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/client/categoryservice"
	"github.com/bearllflee/scholar-track/rpc/achieve/client/propertyservice"
	"github.com/bearllflee/scholar-track/rpc/storage/storage_client"
	"github.com/bearllflee/scholar-track/rpc/system/client/apiservice"
	"github.com/bearllflee/scholar-track/rpc/system/client/casbin"
	"github.com/bearllflee/scholar-track/rpc/system/client/dictionaryservice"
	"github.com/bearllflee/scholar-track/rpc/system/client/menuservice"
	"github.com/bearllflee/scholar-track/rpc/system/client/role"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	User              user.User
	Role              role.Role
	Casbin            casbin.Casbin
	Api               apiservice.ApiService
	Menu              menuservice.MenuService
	Category          achieve.CategoryServiceClient
	Property          achieve.PropertyServiceClient
	Achieve           achieve.AchieveServiceClient
	Storage           storage_client.Storage
	DictionaryService dictionaryservice.DictionaryService
	CasbinRbac        rest.Middleware
	JwtMiddleWare     rest.Middleware
}

func (s *ServiceContext) Enforce(r *http.Request, role string, path string, method string) (bool, error) {
	resp, err := s.Casbin.Enforce(r.Context(), &casbin.EnforceReq{
		Role:   role,
		Path:   path,
		Method: method,
	})
	if err != nil {
		return false, err
	}
	return resp.Result, nil
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:            c,
		User:              user.NewUser(zrpc.MustNewClient(c.System)),
		Role:              role.NewRole(zrpc.MustNewClient(c.System)),
		Casbin:            casbin.NewCasbin(zrpc.MustNewClient(c.System)),
		Api:               apiservice.NewApiService(zrpc.MustNewClient(c.System)),
		Menu:              menuservice.NewMenuService(zrpc.MustNewClient(c.System)),
		Category:          categoryservice.NewCategoryService(zrpc.MustNewClient(c.Achieve)),
		Property:          propertyservice.NewPropertyService(zrpc.MustNewClient(c.Achieve)),
		Achieve:           achieveservice.NewAchieveService(zrpc.MustNewClient(c.Achieve)),
		Storage:           storage_client.NewStorage(zrpc.MustNewClient(c.Storage)),
		DictionaryService: dictionaryservice.NewDictionaryService(zrpc.MustNewClient(c.System)),
	}
	svcCtx.CasbinRbac = middleware.NewCasbinMiddleware(svcCtx).Handle
	svcCtx.JwtMiddleWare = middleware.NewJwtMiddleware(c).Handle
	return svcCtx
}
