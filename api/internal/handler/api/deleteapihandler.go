package api

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/api"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"google.golang.org/grpc/status"
)

func DeleteApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteApiReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := api.NewDeleteApiLogic(r.Context(), svcCtx)
		_, err := l.DeleteApi(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
		} else {
			response.Success(r.Context(), w)
		}
	}
}
