package api

import (
	"errors"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/api"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func UpdateApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateApiReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := api.NewUpdateApiLogic(r.Context(), svcCtx)
		resp, err := l.UpdateApi(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrApiNotFound) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				logx.Error("更新api失败: ", err)
				response.ErrWithMessage(r.Context(), w, "更新api失败")
			}
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
