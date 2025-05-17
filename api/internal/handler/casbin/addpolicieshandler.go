package casbin

import (
	"net/http"

	"github.com/bearllflee/scholar-track/pkg/response"

	"github.com/bearllflee/scholar-track/api/internal/logic/casbin"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddPoliciesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPolicyReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := casbin.NewAddPolicyLogic(r.Context(), svcCtx)
		err := l.AddPolicy(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "添加策略失败")
		} else {
			response.Success(r.Context(), w)
		}
	}
}
