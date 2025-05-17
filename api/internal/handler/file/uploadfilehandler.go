package file

import (
	"io"
	"net/http"
	"strconv"

	filelogic "github.com/bearllflee/scholar-track/api/internal/logic/file"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/response"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		err := r.ParseMultipartForm(5 << 20) // 32MB 限制
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "文件大小超过5MB限制")
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "文件上传失败")
			return
		}
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "文件上传失败")
			return
		}
		req.File = fileBytes
		req.FileName = handler.Filename
		req.FileType = r.FormValue("fileType")
		if req.FileType == "" {
			response.ErrWithMessage(r.Context(), w, "参数错误，缺少文件类型")
			return
		}
		req.BusinessId, err = strconv.ParseUint(r.FormValue("bussinessId"), 10, 64)
		if err != nil || req.BusinessId == 0 {
			response.ErrWithMessage(r.Context(), w, "参数错误，业务ID格式错误")
			return
		}
		if req.BusinessId == 0 {
			response.ErrWithMessage(r.Context(), w, "参数错误，缺少业务ID")
			return
		}
		req.BussinessName = r.FormValue("bussinessName")
		if req.BussinessName == "" {
			response.ErrWithMessage(r.Context(), w, "参数错误，缺少业务名称")
			return
		}
		l := filelogic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "文件上传失败")
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
