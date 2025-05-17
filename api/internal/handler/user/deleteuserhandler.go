package user

import (
	"errors"
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/user"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		_, err := l.DeleteUser(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrUserNotFound) {	
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
			} else {
				logx.Error("删除用户失败: ", err)
				response.ErrWithMessage(r.Context(), w, "删除用户失败")
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
