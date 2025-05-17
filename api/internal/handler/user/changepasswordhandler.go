package user

import (
	"errors"
	"github.com/bearllfleed/scholar-track/api/internal/utils"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"net/http"
	"github.com/bearllfleed/scholar-track/api/internal/logic/user"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}
		req.UserId = utils.GetUserId(r)

		l := user.NewChangePasswordLogic(r.Context(), svcCtx)

		_, err := l.ChangePassword(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrUserNotFound) || errors.Is(err, cerror.ErrPasswordWrong) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				response.ErrWithMessage(r.Context(), w, err.Error())
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
