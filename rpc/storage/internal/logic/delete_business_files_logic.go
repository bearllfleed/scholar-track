package logic

import (
	"context"
	"time"

	"github.com/bearllfleed/scholar-track/rpc/storage/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBusinessFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBusinessFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBusinessFilesLogic {
	return &DeleteBusinessFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBusinessFilesLogic) DeleteBusinessFiles(in *storage.DeleteBusinessFilesRequest) (*storage.FileDeleteResponse, error) {
	files, err := l.svcCtx.FileService.GetFileByBussiness(l.ctx, in.BussnessId, in.BusinessName)
	if err != nil {
		failedDeletion := model.FailedDeletion{
			BusinessID:   in.BussnessId,
			BusinessType: in.BusinessName,
			ErrorReason:  err.Error(),
			RetryCount:   0,
			NextRetryAt:  time.Now().Add(time.Minute * 2),
			ErrorType:    2,
		}
		l.svcCtx.FailedDeletionsService.CreateFailedDeletions(l.ctx, &failedDeletion)
	}
	for _, file := range files {
		err = l.svcCtx.FileService.DeleteFile(l.ctx, file.Id)
		if err != nil {
			failedDeletion := model.FailedDeletion{
				FileID:       file.Id,
				FileKey:      file.FileKey,
				BusinessID:   in.BussnessId,
				BusinessType: in.BusinessName,
				ErrorReason:  err.Error(),
				RetryCount:   0,
				NextRetryAt:  time.Now().Add(time.Minute * 2),
				ErrorType:    0,
			}
			l.svcCtx.FailedDeletionsService.CreateFailedDeletions(l.ctx, &failedDeletion)
		}
		err = l.svcCtx.StorageService.DeleteFile(l.ctx, file.FileKey)
		if err != nil {
			failedDeletion := model.FailedDeletion{
				FileID:       file.Id,
				FileKey:      file.FileKey,
				BusinessID:   in.BussnessId,
				BusinessType: in.BusinessName,
				ErrorReason:  err.Error(),
				RetryCount:   0,
				NextRetryAt:  time.Now().Add(time.Minute * 2),
				ErrorType:    1,
			}
			l.svcCtx.FailedDeletionsService.CreateFailedDeletions(l.ctx, &failedDeletion)
		}
	}
	return &storage.FileDeleteResponse{}, nil
}
