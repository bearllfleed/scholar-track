package menu

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/menu"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func CreateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := menu.NewCreateMenuLogic(r.Context(), svcCtx)
		_, err := l.CreateMenu(&req)
		if err != nil {
			if cerror.CheckIsDefinedMenuError(err) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				logx.Error("创建菜单失败: ", err)
				response.ErrWithMessage(r.Context(), w, "创建菜单失败")
			}
			return
		}
		response.Success(r.Context(), w)
	}
}

