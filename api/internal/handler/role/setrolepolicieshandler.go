package role

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/role"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetRolePoliciesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetRolePoliciesReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := role.NewSetRolePoliciesLogic(r.Context(), svcCtx)
		_, err := l.SetRolePolicies(&req)
		if err != nil {
			logx.Error("更新角色权限失败: ", err)
			response.ErrWithMessage(r.Context(), w, "更新角色权限失败")
		} else {
			response.Success(r.Context(), w)
		}
	}
}
