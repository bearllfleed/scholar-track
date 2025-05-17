package casbin

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterCasbinRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	// Casbin routes require JWT and Casbin middleware
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleWare, serverCtx.CasbinRbac},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/casbin/addPolicies",
					Handler: AddPoliciesHandler(serverCtx),
				},
			}...,
		),
	)
}
