package category

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/achievement/category"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func QueryCategoryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCategoryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := category.NewQueryCategoryDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryCategoryDetail(&req)
		if err != nil {
			if cerror.CheckIsDefinedCategoryError(err) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
				return
			} else {
				l.Logger.Error("查询分类详情失败: ", err)
				response.ErrWithMessage(r.Context(), w, "查询分类详情失败")
				return
			}
		}
		response.SuccessWithData(r.Context(), w, resp)
	}
}
