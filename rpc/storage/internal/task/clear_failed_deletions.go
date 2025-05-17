package task

import (
	"context"
	"math"
	"time"

	"github.com/bearllfleed/scholar-track/rpc/storage/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/storage/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClearFailedDeletionsTask struct {
	svcCtx *svc.ServiceContext
}

func NewClearFailedDeletionsTask(svcCtx *svc.ServiceContext) *ClearFailedDeletionsTask {
	return &ClearFailedDeletionsTask{
		svcCtx: svcCtx,
	}
}

func (t *ClearFailedDeletionsTask) Start(ctx context.Context) {
	logx.Info("ClearFailedDeletionsTask started")
	ticker := time.NewTicker(time.Duration(t.svcCtx.Config.FailedDeletion.Interval) * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logx.Info("ClearFailedDeletionsTask stopped")
			return
		case <-ticker.C:
			t.processBatch(ctx)
		}
	}
}

func (t *ClearFailedDeletionsTask) processBatch(ctx context.Context) {
	// 1. 查询待处理记录
	records, err := t.svcCtx.FailedDeletionsService.QueryWaitToHandle(ctx, t.svcCtx.Config.FailedDeletion.BatchSize)
	if err != nil {
		logx.Errorf("QueryWaitToHandle error: %v", err)
		return
	}

	// 2. 批量删除文件
	for _, record := range records {
		if record.ErrorType == 0 {
			if err := t.svcCtx.FileService.DeleteFile(ctx, record.FileID); err != nil {
				// 更新重试信息
				if err := t.handleFailedDeletion(ctx, record, err); err != nil {
					logx.Errorf("Update failed deletion record error: %v", err)
				}
				continue
			}
		} else if record.ErrorType == 1 {
			if err := t.svcCtx.StorageService.DeleteFile(ctx, record.FileKey); err != nil {
				if err := t.handleFailedDeletion(ctx, record, err); err != nil {
					logx.Errorf("Update failed deletion record error: %v", err)
				}
				continue
			}
		}

		// 删除成功，移除记录
		if err := t.svcCtx.FailedDeletionsService.DeleteFailedDeletion(ctx, record.Id); err != nil {
			logx.Errorf("Delete failed deletion record error: %v", err)
		}
	}
}

func (t *ClearFailedDeletionsTask) handleFailedDeletion(ctx context.Context, record *model.FailedDeletion, err error) error {
	// 达到最大重试次数
	if record.RetryCount >= t.svcCtx.Config.FailedDeletion.MaxRetry-1 {
		logx.Errorf("File deletion failed after max retries, fileKey: %s, error: %v", record.FileKey, err)
		return t.svcCtx.FailedDeletionsService.DeleteFailedDeletion(ctx, record.Id) // 或标记为最终失败
	}

	// 更新重试信息
	record.RetryCount++
	record.NextRetryAt = time.Now().Add(time.Duration(math.Pow(2, float64(record.RetryCount))) * time.Minute)
	record.ErrorReason = err.Error()
	return t.svcCtx.FailedDeletionsService.UpdateFailedDeletion(ctx, record)
}
