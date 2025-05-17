package api

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/api"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryAllApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryAllApiReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := api.NewQueryAllApiLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllApi(&req)
		if err != nil {
			l.Logger.Error("查询api失败: ", err)
			response.ErrWithMessage(r.Context(), w, "查询api失败")
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
