package role

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterRoleRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodDelete,
				Path:    "/role/:id",
				Handler: DeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/add",
				Handler: AddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/queryRolePolicies/:roleId",
				Handler: QueryRolePoliciesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/setRoleInfo",
				Handler: SetRoleInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/setRolePolicies",
				Handler: SetRolePoliciesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/tree",
				Handler: RoleTreeHandler(serverCtx),
			},
		},
	)
}
