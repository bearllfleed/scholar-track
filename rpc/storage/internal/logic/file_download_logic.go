package logic

import (
	"context"
	"io"

	"github.com/bearllflee/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDownloadLogic) FileDownload(in *storage.FileDownloadRequest) (*storage.FileDownloadResponse, error) {
	file, err := l.svcCtx.FileService.GetFileById(l.ctx, in.FileId)
	if err != nil {
		return nil, err
	}
	fileReader, err := l.svcCtx.StorageService.DownloadFile(l.ctx, file.FileKey)
	if err != nil {
		return nil, err
	}
	fileData, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, err
	}
	return &storage.FileDownloadResponse{
		FileData: fileData,
		FileName: file.FileName,
		FileType: file.FileType,
		FileSize: uint64(file.FileSize),
	}, nil
}
