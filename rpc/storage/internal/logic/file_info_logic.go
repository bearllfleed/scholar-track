package logic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileInfoLogic) FileInfo(in *storage.FileInfoRequest) (*storage.FileInfoResponse, error) {
	file, err := l.svcCtx.FileService.GetFileById(l.ctx, in.FileId)
	if err != nil {
		return nil, err
	}
	return &storage.FileInfoResponse{
		FileId:      file.Id,
		FileUrl:     file.FileKey,
		FileName:    file.FileName,
		FileType:    file.FileType,
		FileSize:    uint64(file.FileSize),
		BussinessId: file.BusinessId,
		CreatedAt:   file.CreatedAt.Unix(),
		UpdatedAt:   file.UpdatedAt.Unix(),
		DeletedAt:   file.DeletedAt.Time.Unix(),
	}, nil
}
