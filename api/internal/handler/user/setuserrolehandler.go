package user

import (
	"errors"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/user"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func SetUserRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := user.NewSetUserRoleLogic(r.Context(), svcCtx)
		_, err := l.SetUserRole(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrInvalidRole) || errors.Is(err, cerror.ErrUserNotFound) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("设置用户角色失败: ", err)
				response.ErrWithMessage(r.Context(), w, "设置用户角色失败")
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
