package response

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(ctx context.Context, w http.ResponseWriter, code int, data interface{}, message string) {
	resp := &Response{
		Code:    code,
		Data:    data,
		Message: message,
	}
	httpx.OkJsonCtx(ctx, w, resp)
}

func Error(ctx context.Context, w http.ResponseWriter) {
	Result(ctx, w, ERROR, nil, "操作失败")
}

func ErrWithMessage(ctx context.Context, w http.ResponseWriter, message string) {
	Result(ctx, w, ERROR, nil, message)
}

func Success(ctx context.Context, w http.ResponseWriter) {
	Result(ctx, w, SUCCESS, nil, "操作成功")
}

func SuccessWithData(ctx context.Context, w http.ResponseWriter, data interface{}) {
	Result(ctx, w, SUCCESS, data, "查询成功")
}

func SuccessWithDetail(ctx context.Context, w http.ResponseWriter, data interface{}, message string) {
	Result(ctx, w, SUCCESS, data, message)
}

func ErrWithNoAuth(ctx context.Context, w http.ResponseWriter, message string) {
	Result(ctx, w, ERROR, nil, message)
}

func SuccessWithFile(ctx context.Context, w http.ResponseWriter, fileData []byte, fileName string, fileType string, fileSize int64) {
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", convertFileTypeToContentType(fileType))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))
	w.WriteHeader(http.StatusOK)
	w.Write(fileData)
}

func convertFileTypeToContentType(fileType string) string {
	switch fileType {
	case "pdf":
		return "application/pdf"
	case "doc":
		return "application/msword"
	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case "xls":
		return "application/vnd.ms-excel"
	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case "ppt":
		return "application/vnd.ms-powerpoint"
	case "pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case "txt":
		return "text/plain"
	case "jpg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	case "bmp":
		return "image/bmp"
	case "ico":
		return "image/x-icon"
	case "mp4":
		return "video/mp4"
	case "mp3":
		return "audio/mpeg"
	case "wav":
		return "audio/x-wav"
	case "avi":
		return "video/x-msvideo"
	case "mov":
		return "video/quicktime"
	case "zip":
		return "application/zip"
	case "rar":
		return "application/x-rar-compressed"
	case "7z":
		return "application/x-7z-compressed"
	}
	return "application/octet-stream"
}

