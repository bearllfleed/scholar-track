package model

import (
	"time"

	"github.com/bearllfleed/scholar-track/pkg/global"
)

type FailedDeletion struct {
	global.StModel
	FileID       uint64    `json:"fileId" gorm:"type:bigint unsigned;column:file_id;comment:文件ID"`
	FileKey      string    `json:"fileKey" gorm:"type:varchar(255);column:file_key;comment:文件key"`
	BusinessID   uint64    `json:"businessId" gorm:"type:bigint unsigned;column:business_id;comment:关联业务ID"`
	BusinessType string    `json:"businessType" gorm:"type:varchar(50);column:business_type;comment:业务类型"`
	ErrorReason  string    `json:"errorReason" gorm:"type:text;column:error_reason;comment:错误原因"`
	RetryCount   int       `json:"retryCount" gorm:"type:int;column:retry_count;default:0;comment:重试次数"`
	NextRetryAt  time.Time `json:"nextRetryAt" gorm:"type:timestamp;column:next_retry_at;comment:下次重试时间"`
	ErrorType    int       `json:"errorType" gorm:"type:int;column:error_type;default:0;comment:错误类型(0文件表 1minio文件存储 2 全部)"`
}

func (FailedDeletion) TableName() string {
	return "st_failed_deletions"
}
