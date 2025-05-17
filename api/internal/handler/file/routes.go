package file

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterFileRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodDelete,
				Path:    "/file/:id",
				Handler: DeleteFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/:id",
				Handler: QueryFileDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/bussiness",
				Handler: GetBussinessFilesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/download/:id",
				Handler: DownloadFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: UploadFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/url",
				Handler: GetFileUrlHandler(serverCtx),
			},
		},
	)
}
