package user

import (
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/user"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}
		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			if cerror.CheckIsDefinedUserError(err) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
				return
			}
			l.Logger.Error("注册失败: ", err)
			response.ErrWithMessage(r.Context(), w, "注册失败")
			return
		}
		response.SuccessWithData(r.Context(), w, resp)
	}
}
