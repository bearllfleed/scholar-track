package user

import (
	"github.com/bearllflee/scholar-track/api/internal/utils"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/user"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QuerySelfInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySelfInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误:", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}
		l := user.NewQuerySelfInfoLogic(r.Context(), svcCtx)
		req.Id = utils.GetUserId(r)
		resp, err := l.QuerySelfInfo(&req)
		if err != nil {
			l.Logger.Error("查询失败: ", err)
			response.ErrWithMessage(r.Context(), w, err.Error())
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
