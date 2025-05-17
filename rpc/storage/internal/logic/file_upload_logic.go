package logic

import (
	"bytes"
	"context"
	"io"

	"github.com/bearllfleed/scholar-track/rpc/storage/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadLogic) FileUpload(in *storage.FileUploadRequest) (*storage.FileUploadResponse, error) {
	fileReader := io.NopCloser(bytes.NewReader(in.FileData))
	objectName, err := l.svcCtx.StorageService.UploadFile(l.ctx, fileReader, in.FileName, in.BussinessId, in.BussinessName, int64(len(in.FileData)), in.FileType)
	if err != nil {
		return nil, err
	}
	file := &model.File{
		FileName:      in.FileName,
		FileKey:       objectName,
		FileType:      in.FileType,
		FileSize:      int64(len(in.FileData)),
		BusinessId:    in.BussinessId,
		BussinessName: in.BussinessName,
	}
	file, err = l.svcCtx.FileService.UploadFile(l.ctx, file)
	if err != nil {
		l.svcCtx.StorageService.DeleteFile(l.ctx, objectName)
		return nil, err
	}
	return &storage.FileUploadResponse{
		FileId: file.Id,
	}, nil
}
