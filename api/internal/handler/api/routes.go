package api

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterApiRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	// All API routes require JWT and Casbin middleware
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleWare, serverCtx.CasbinRbac},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api",
					Handler: CreateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api",
					Handler: QueryAllApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api",
					Handler: UpdateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/:id",
					Handler: DeleteApiHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/:id",
					Handler: QueryApiDetailHandler(serverCtx),
				},
			}...,
		),
	)
}
