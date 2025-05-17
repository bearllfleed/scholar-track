package service

import (
	"errors"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"gorm.io/gorm"
)

var ApiServiceApp = new(ApiService)

type ApiService struct{}

func (s *ApiService) DeleteApi(id uint64) error {
	err := global.DB.Model(&model.Api{}).Where("id = ?", id).Delete(&model.Api{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ApiService) UpdateApi(api *model.Api) (*model.Api, error) {
	err := global.DB.Save(api).Error
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (s *ApiService) QueryApiDetail(id uint64) (*model.Api, error) {
	var api model.Api
	err := global.DB.Model(&model.Api{}).Where("id = ?", id).First(&api).Error
	if err != nil {
		return nil, err
	}
	return &api, nil
}

func (s *ApiService) QueryAllApi() ([]*model.Api, error) {
	var apis []*model.Api
	err := global.DB.Model(&model.Api{}).Order("api_group DESC").Find(&apis).Error
	if err != nil {
		return nil, err
	}
	return apis, nil
}

func (s *ApiService) CreateApi(api *model.Api) error {
	return global.DB.Create(api).Error
}
func (s *ApiService) CheckApiExist(apiId uint64, path string, method string) (bool, error) {
	var api model.Api
	err := global.DB.Model(&model.Api{}).Where("path = ? AND method = ?", path, method).First(&api).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return api.Id != apiId, nil
}

func (s *ApiService) QueryApiByPathAndMethod(path string, method string) (*model.Api, error) {
	var api model.Api
	err := global.DB.Model(&model.Api{}).Where("path = ? AND method = ?", path, method).First(&api).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.ErrApiNotFound
		}
		return nil, err
	}
	return &api, nil
}
