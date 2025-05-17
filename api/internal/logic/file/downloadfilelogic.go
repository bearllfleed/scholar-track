package file

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/storage/storage_client"
	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.DownloadFileReq) (resp *types.DownloadFileResp, err error) {
	fileData, err := l.svcCtx.Storage.FileDownload(l.ctx, &storage_client.FileDownloadRequest{FileId: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DownloadFileResp{FileData: fileData.FileData, FileName: fileData.FileName, FileType: fileData.FileType, FileSize: int64(fileData.FileSize)}, nil
}
