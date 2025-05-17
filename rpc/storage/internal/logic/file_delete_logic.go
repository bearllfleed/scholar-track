package logic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeleteLogic) FileDelete(in *storage.FileDeleteRequest) (*storage.FileDeleteResponse, error) {
	file, err := l.svcCtx.FileService.GetFileById(l.ctx, in.FileId)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.FileService.DeleteFile(l.ctx, file.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.StorageService.DeleteFile(l.ctx, file.FileKey)
	if err != nil {
		return nil, err
	}
	return &storage.FileDeleteResponse{}, nil
}
