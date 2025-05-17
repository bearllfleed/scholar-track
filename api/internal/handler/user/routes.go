package user

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterUserRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	// Public routes (no middleware)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: RegisterHandler(serverCtx),
			},
		},
	)

	// Protected routes (require JWT and Casbin)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleWare, serverCtx.CasbinRbac},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/:id",
					Handler: QueryUserDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/:id",
					Handler: DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/changePassword",
					Handler: ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: QueryUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/resetPassword",
					Handler: ResetPasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/selfInfo",
					Handler: QuerySelfInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/setSelfInfo",
					Handler: SetSelfInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/setUserInfo",
					Handler: SetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/setUserRole",
					Handler: SetUserRoleHandler(serverCtx),
				},
			}..., // Note the ellipsis for variadic arguments
		),
	)
}
