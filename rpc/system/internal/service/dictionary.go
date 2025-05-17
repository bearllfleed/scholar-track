package service

import (
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
)

var DictionaryServiceApp = new(DictionaryService)

type DictionaryService struct {
}

func (s *DictionaryService) CreateDictionary(dictionary *model.Dictionary) error {
	return global.DB.Create(dictionary).Error
}

func (s *DictionaryService) DeleteDictionary(id uint64) error {
	return global.DB.Delete(&model.Dictionary{}, id).Error
}

func (s *DictionaryService) UpdateDictionary(dictionary *model.Dictionary) (*model.Dictionary, error) {
	err := global.DB.Save(dictionary).Error
	if err != nil {
		return nil, err
	}
	return dictionary, nil
}

func (s *DictionaryService) QueryDictionaryDetail(id uint64) (*model.Dictionary, error) {
	var dictionary model.Dictionary
	err := global.DB.Preload("DictionaryDetails").First(&dictionary, id).Error
	if err != nil {
		return nil, err
	}
	return &dictionary, nil
}

func (s *DictionaryService) QueryAllDictionary() ([]*model.Dictionary, error) {
	var dictionaries []*model.Dictionary
	err := global.DB.Preload("DictionaryDetails").Find(&dictionaries).Error
	if err != nil {
		return nil, err
	}
	return dictionaries, nil
}

func (s *DictionaryService) QueryDictionaryList(page, pageSize int64, name, dictionaryType string) ([]*model.Dictionary, int64, error) {
	var dictionaries []*model.Dictionary
	var total int64

	db := global.DB.Model(&model.Dictionary{})

	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	if dictionaryType != "" {
		db = db.Where("type = ?", dictionaryType)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Preload("DictionaryDetails").Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&dictionaries).Error
	if err != nil {
		return nil, 0, err
	}

	return dictionaries, total, nil
}
