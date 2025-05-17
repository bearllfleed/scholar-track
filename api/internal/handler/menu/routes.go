package menu

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterMenuRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/menu",
				Handler: CreateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/menu",
				Handler: UpdateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/:id",
				Handler: QueryMenuDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu/:id",
				Handler: DeleteMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/all",
				Handler: QueryAllMenuTreeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/role",
				Handler: UpdateRoleMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/role/:roleId",
				Handler: QueryRoleMenuTreeHandler(serverCtx),
			},
		},
	)
}
