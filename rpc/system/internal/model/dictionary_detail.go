package model

import "github.com/bearllflee/scholar-track/rpc/system/internal/global"

type DictionaryDetail struct {
	global.StModel
	Label        string `json:"label" form:"label" gorm:"column:label;comment:展示值"`                        // 展示值
	Value        string `json:"value" form:"value" gorm:"column:value;comment:字典值"`                        // 字典值
	Extend       string `json:"extend" form:"extend" gorm:"column:extend;comment:扩展值"`                     // 扩展值
	Status       *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`                    // 启用状态
	Sort         int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`                          // 排序标记
	DictionaryId uint64 `json:"dictionaryID" form:"dictionaryID" gorm:"column:dictionary_id;comment:关联标记"` // 关联标记
}

func (DictionaryDetail) TableName() string {
	return "st_dictionary_details"
}
