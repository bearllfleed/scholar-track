package role

import (
	"errors"
	"net/http"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"google.golang.org/grpc/status"

	"github.com/bearllfleed/scholar-track/api/internal/logic/role"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := role.NewAddRoleLogic(r.Context(), svcCtx)
		resp, err := l.AddRole(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrRoleHasExists) || errors.Is(err, cerror.ErrParentRoleNotExists) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("添加角色失败: ", err)
				response.ErrWithMessage(r.Context(), w, "添加角色失败")
			}
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
