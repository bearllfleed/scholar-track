package achievement

import (
	"net/http"

	achievementbasic "github.com/bearllflee/scholar-track/api/internal/handler/achievement/basic"
	achievementcategory "github.com/bearllflee/scholar-track/api/internal/handler/achievement/category"
	achievementproperty "github.com/bearllflee/scholar-track/api/internal/handler/achievement/property"
	"github.com/bearllflee/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterAchievementRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	// All achievement routes require JWT and Casbin middleware
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleWare, serverCtx.CasbinRbac},
			[]rest.Route{
				// Basic Achievement
				{
					Method:  http.MethodPost,
					Path:    "/achievement",
					Handler: achievementbasic.UploadAchievementHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/achievement/:id",
					Handler: achievementbasic.DeleteAchievementHandler(serverCtx),
				},
				// Category
				{
					Method:  http.MethodGet,
					Path:    "/category",
					Handler: achievementcategory.QueryCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category",
					Handler: achievementcategory.AddCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/category",
					Handler: achievementcategory.UpdateCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/category/:id",
					Handler: achievementcategory.DeleteCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/category/:id",
					Handler: achievementcategory.QueryCategoryDetailHandler(serverCtx),
				},
				// Property
				{
					Method:  http.MethodGet,
					Path:    "/property",
					Handler: achievementproperty.QueryPropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property",
					Handler: achievementproperty.AddPropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/property",
					Handler: achievementproperty.UpdatePropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/property/:id",
					Handler: achievementproperty.DeletePropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/property/:id",
					Handler: achievementproperty.QueryPropertyDetailHandler(serverCtx),
				},
			}..., // Note the ellipsis for variadic arguments
		),
	)
}
