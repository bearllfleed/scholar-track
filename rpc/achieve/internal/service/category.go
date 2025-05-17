package service

import (
	"context"
	"errors"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/model"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (c *CategoryService) UpdateCategory(ctx context.Context, category *model.Category) error {
	err := c.db.WithContext(ctx).Model(&model.Category{}).Where("id = ?", category.Id).Updates(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryService) QueryCategoryDetail(ctx context.Context, id uint64) (*model.Category, error) {
	var category model.Category
	err := c.db.WithContext(ctx).Model(&model.Category{}).Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.ErrCategoryNotFound
		}
		return nil, err
	}
	return &category, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, in *achieve.DeleteCategoryReq) error {
	err := c.db.WithContext(ctx).Model(&model.Category{}).Where("id = ?", in.Id).Delete(&model.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryService) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	err := c.db.WithContext(ctx).Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) QueryCategoryList(ctx context.Context, name string, typeStr string, status int32, page int, pageSize int) (error, int64, []*model.Category) {
	var categories []*model.Category
	var total int64
	var db = c.db.WithContext(ctx)
	if name != "" {
		err := db.Model(&model.Category{}).Where("name LIKE ?", "%"+name+"%").Where("type = ?", typeStr).Where("status = ?", status).Count(&total).Error
		if err != nil {
			return err, 0, nil
		}
	}
	if typeStr != "" {
		err := db.Model(&model.Category{}).Where("type = ?", typeStr).Where("status = ?", status).Count(&total).Error
		if err != nil {
			return err, 0, nil
		}
	}
	if status != 0 {
		err := db.Model(&model.Category{}).Where("status = ?", status).Count(&total).Error
		if err != nil {
			return err, 0, nil
		}
	}
	return nil, total, categories
}
