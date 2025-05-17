package menu

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/menu"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func QueryMenuDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryMenuDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := menu.NewQueryMenuDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryMenuDetail(&req)
		if err != nil {
			if cerror.CheckIsDefinedMenuError(err) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				logx.Error("查询菜单详情失败: ", err)
				response.ErrWithMessage(r.Context(), w, "查询菜单详情失败")
			}
			return
		}
		response.SuccessWithData(r.Context(), w, resp)
	}
}
