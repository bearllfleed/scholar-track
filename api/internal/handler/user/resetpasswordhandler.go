package user

import (
	"errors"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/user"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ResetPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err.Error())
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := user.NewResetPasswordLogic(r.Context(), svcCtx)
		_, err := l.ResetPassword(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrUserNotFound) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("重置失败: ", err.Error())
				response.ErrWithMessage(r.Context(), w, "重置失败")
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
