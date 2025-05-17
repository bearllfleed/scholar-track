package dictionary_detail

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterDictionaryDetailRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	// All dictionary detail routes require JWT and Casbin middleware
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleWare, serverCtx.CasbinRbac},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/detail",
					Handler: CreateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/dictionary/detail",
					Handler: UpdateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/dictionary/detail/:id",
					Handler: DeleteDictionaryDetailHandler(serverCtx),
				},
			}...,
		),
	)
}
