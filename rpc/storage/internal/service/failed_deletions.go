package service

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/storage/internal/model"
	"gorm.io/gorm"
)

type FailedDeletionsService struct {
	db *gorm.DB
}

func NewFailedDeletionsService(db *gorm.DB) *FailedDeletionsService {
	return &FailedDeletionsService{db: db}
}

func (s *FailedDeletionsService) CreateFailedDeletions(ctx context.Context, failedDeletions *model.FailedDeletion) error {
	return s.db.Model(&model.FailedDeletion{}).Create(failedDeletions).Error
}

func (s *FailedDeletionsService) QueryFailedDeletionsByBusiness(ctx context.Context, businessName string, businessId uint64) (failedDeletions []*model.FailedDeletion, err error) {
	err = s.db.WithContext(ctx).Model(&model.FailedDeletion{}).Find(&failedDeletions).Error
	return
}

func (s *FailedDeletionsService) DeleteFailedDeletion(ctx context.Context, id uint64) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&model.FailedDeletion{}).Error
}

func (s *FailedDeletionsService) UpdateFailedDeletion(ctx context.Context, record *model.FailedDeletion) error {
	return s.db.WithContext(ctx).Model(record).Updates(map[string]interface{}{
		"retry_count":   record.RetryCount,
		"next_retry_at": record.NextRetryAt,
		"error_reason":  record.ErrorReason,
	}).Error
}

func (s *FailedDeletionsService) QueryWaitToHandle(ctx context.Context, limit int) (failedDeletions []*model.FailedDeletion, err error) {
	err = s.db.Model(&model.FailedDeletion{}).Where("next_retry_at <= NOW() AND retry_count < ?", 3).Limit(limit).Find(&failedDeletions).Error
	return
}
