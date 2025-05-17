package basic

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/achievement/basic"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteAchievementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteAchievementReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := basic.NewDeleteAchievementLogic(r.Context(), svcCtx)
		_, err := l.DeleteAchievement(&req)
		if err != nil {
			l.Logger.Error("删除成果失败: ", err)
			response.ErrWithMessage(r.Context(), w, "删除成果失败")
		} else {
			response.Success(r.Context(), w)
		}
	}
}
