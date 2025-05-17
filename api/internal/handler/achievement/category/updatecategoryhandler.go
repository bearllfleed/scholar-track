package category

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/achievement/category"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func UpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := category.NewUpdateCategoryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCategory(&req)
		if err != nil {
			if cerror.CheckIsDefinedCategoryError(err) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
				return
			} else {
				l.Logger.Error("更新分类失败: ", err)
				response.ErrWithMessage(r.Context(), w, "更新分类失败")
				return
			}
		}
		response.SuccessWithData(r.Context(), w, resp)
	}
}
