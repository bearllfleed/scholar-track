package logic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBussinessFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBussinessFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBussinessFilesLogic {
	return &GetBussinessFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBussinessFilesLogic) GetBussinessFiles(in *storage.GetBussinessFilesRequest) (*storage.GetBussinessFilesResponse, error) {
	files, err := l.svcCtx.FileService.GetFileByBussiness(l.ctx, in.BussinessId, in.BussinessName)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		file.FileUrl, err = l.svcCtx.StorageService.GetFileUrl(l.ctx, file.FileKey)
		if err != nil {
			return nil, err
		}
	}

	filesResponse := make([]*storage.FileInfoResponse, 0)
	for _, file := range files {
		filesResponse = append(filesResponse, &storage.FileInfoResponse{
			FileId:      file.Id,
			FileKey:     file.FileKey,
			FileUrl:     file.FileUrl,
			FileName:    file.FileName,
			FileType:    file.FileType,
			FileSize:    uint64(file.FileSize),
			BussinessId: file.BusinessId,
			CreatedAt:   file.CreatedAt.Unix(),
			UpdatedAt:   file.UpdatedAt.Unix(),
			DeletedAt:   file.DeletedAt.Time.Unix(),
		})
	}

	return &storage.GetBussinessFilesResponse{
		Files: filesResponse,
	}, nil
}
