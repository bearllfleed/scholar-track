package role

import (
	"errors"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/role"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误: ", err)
			response.ErrWithMessage(r.Context(),w, err.Error())
			return
		}

		l := role.NewDeleteRoleLogic(r.Context(), svcCtx)
		_, err := l.DeleteRole(&req)
		if err != nil {
			if errors.Is(err, cerror.ErrRoleNotExists) {
				response.ErrWithMessage(r.Context(), w, "角色不存在")
			} else {
				l.Logger.Error("删除角色失败: ", err)
				response.ErrWithMessage(r.Context(), w, err.Error())
			}
		} else {
			response.Success(r.Context(), w)
		}
	}
}
