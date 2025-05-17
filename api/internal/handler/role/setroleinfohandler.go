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

func SetRoleInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetRoleInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := role.NewSetRoleInfoLogic(r.Context(), svcCtx)
		_, err := l.SetRoleInfo(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrRoleNotExists) || errors.Is(err, cerror.ErrRoleHasExists) || errors.Is(err, cerror.ErrParentRoleNotExists) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("设置角色信息失败: ", err)
				response.ErrWithMessage(r.Context(), w, "设置角色信息失败")
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
