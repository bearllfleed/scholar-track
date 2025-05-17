package file

import (
	"net/http"

	filelogic "github.com/bearllflee/scholar-track/api/internal/logic/file"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBussinessFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBussinessFilesReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(r.Context(), "获取业务文件失败", err)
			response.ErrWithMessage(r.Context(), w, "获取业务文件失败")
			return
		}

		l := filelogic.NewGetBussinessFilesLogic(r.Context(), svcCtx)
		resp, err := l.GetBussinessFiles(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "获取业务文件失败")
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
