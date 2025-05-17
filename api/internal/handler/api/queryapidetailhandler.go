package api

import (
	"errors"
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/api"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"google.golang.org/grpc/status"

	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryApiDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryApiDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := api.NewQueryApiDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryApiDetail(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrApiNotFound) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("查询api详情失败: ", err)
				response.ErrWithMessage(r.Context(), w, "查询api详情失败")
			}
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
