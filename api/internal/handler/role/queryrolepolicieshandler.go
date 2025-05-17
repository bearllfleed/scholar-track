package role

import (
	"errors"
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/role"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func QueryRolePoliciesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryRolePoliciesReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error( "参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := role.NewQueryRolePoliciesLogic(r.Context(), svcCtx)
		resp, err := l.QueryRolePolicies(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrRoleNotExists) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("获取角色权限失败: ", err)
				response.ErrWithMessage(r.Context(), w, "获取角色权限失败")
			}
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
