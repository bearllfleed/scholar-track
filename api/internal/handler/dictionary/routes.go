package dictionary

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterDictionaryRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	// All dictionary routes require JWT and Casbin middleware
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleWare, serverCtx.CasbinRbac},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary",
					Handler: CreateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/dictionary",
					Handler: UpdateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/dictionary/:id",
					Handler: DeleteDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/dictionary/:id",
					Handler: QueryDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/list",
					Handler: QueryDictionaryListHandler(serverCtx),
				},
			}...,
		),
	)
}
