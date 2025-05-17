package user

import (
	"errors"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/user"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrUserNotFound) || errors.Is(err, cerror.ErrPasswordWrong) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				l.Logger.Error("登录失败: ", err)
				response.ErrWithMessage(r.Context(), w, "登录失败")
			}
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
