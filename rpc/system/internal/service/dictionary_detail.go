package service

import (
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
)

var DictionaryDetailServiceApp = new(DictionaryDetailService)

type DictionaryDetailService struct {
}

func (s *DictionaryDetailService) CreateDictionaryDetail(dictionaryDetail *model.DictionaryDetail) error {
	return global.DB.Create(dictionaryDetail).Error
}

func (s *DictionaryDetailService) DeleteDictionaryDetail(id uint64) error {
	return global.DB.Delete(&model.DictionaryDetail{}, id).Error
}

func (s *DictionaryDetailService) UpdateDictionaryDetail(dictionaryDetail *model.DictionaryDetail) (*model.DictionaryDetail, error) {
	err := global.DB.Save(dictionaryDetail).Error
	if err != nil {
		return nil, err
	}
	return dictionaryDetail, nil
}

func (s *DictionaryDetailService) QueryDictionaryDetailDetail(id uint64) (*model.DictionaryDetail, error) {
	var dictionaryDetail model.DictionaryDetail
	err := global.DB.First(&dictionaryDetail, id).Error
	if err != nil {
		return nil, err
	}
	return &dictionaryDetail, nil
}

func (s *DictionaryDetailService) QueryAllDictionaryDetail() ([]*model.DictionaryDetail, error) {
	var dictionaryDetails []*model.DictionaryDetail
	err := global.DB.Find(&dictionaryDetails).Error
	if err != nil {
		return nil, err
	}
	return dictionaryDetails, nil
}
