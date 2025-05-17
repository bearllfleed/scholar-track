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
)

func CreateApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateApiReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := api.NewCreateApiLogic(r.Context(), svcCtx)
		resp, err := l.CreateApi(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrApiAlreadyExists) {
				response.ErrWithMessage(r.Context(), w, "api已存在")
			} else {
				l.Logger.Error("创建api失败: ", err)
				response.ErrWithMessage(r.Context(), w, "创建api失败")
			}
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
