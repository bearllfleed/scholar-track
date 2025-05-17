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

func AddCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := category.NewAddCategoryLogic(r.Context(), svcCtx)
		if req.Status == 0 {
			req.Status = 1
		}
		resp, err := l.AddCategory(&req)
		if err != nil {
			if cerror.CheckIsDefinedCategoryError(err) {
				response.ErrWithMessage(r.Context(), w, status.Convert(err).Message())
				return
			} else {
				l.Logger.Error("添加分类失败: ", err)
				response.ErrWithMessage(r.Context(), w, "添加分类失败")
				return
			}
		}
		response.SuccessWithData(r.Context(), w, resp)
	}
}
