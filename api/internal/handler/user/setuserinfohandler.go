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

func SetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := user.NewSetUserInfoLogic(r.Context(), svcCtx)
		_, err := l.SetUserInfo(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrUserNotFound) || errors.Is(err, cerror.ErrUserHasExists) || errors.Is(err, cerror.ErrPhoneHasExists) || errors.Is(err, cerror.ErrEmailHasExists) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("更新失败: ", err)
				response.ErrWithMessage(r.Context(), w, err.Error())
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
