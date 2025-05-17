package file

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/file"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteFileReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, err.Error())
			return
		}

		l := file.NewDeleteFileLogic(r.Context(), svcCtx)
		_, err := l.DeleteFile(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "删除文件失败")
		} else {
			response.Success(r.Context(), w)
		}
	}
}
