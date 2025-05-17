package service

import (
	"context"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/model"
	"gorm.io/gorm"
)

type AchieveService struct {
	db *gorm.DB
}

func NewAchieveService(db *gorm.DB) *AchieveService {
	return &AchieveService{db: db}
}

func (s *AchieveService) UploadAchieve(ctx context.Context, achieve *model.AchieveBasic) (*model.AchieveBasic, error) {
	var achieveBasic = &model.AchieveBasic{}

	s.db.WithContext(ctx).Model(&model.AchieveBasic{}).Where("code = ?", achieve.Code).First(achieveBasic)
	// 如果存在，则返回错误
	if achieveBasic.Id != 0 {
		return nil, cerror.ErrAchieveHasExists
	}
	id, err := global.Snowflake.NextID()
	if err != nil {
		return nil, cerror.ErrGenerateSnowflakeIdFailed
	}
	achieve.Id = uint64(id)
	// 创建新记录
	if err := s.db.WithContext(ctx).Create(achieve).Error; err != nil {
		return nil, err
	}
	return achieve, nil
}

func (s *AchieveService) DeleteAchieve(ctx context.Context, id uint64) error {
	db := s.db.WithContext(ctx).Model(&model.AchieveBasic{})
	// 物理删除
	if err := db.Delete(&model.AchieveBasic{}, id).Error; err != nil {
		return err
	}
	return nil
}
