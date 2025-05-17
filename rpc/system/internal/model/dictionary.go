package model

import "github.com/bearllflee/scholar-track/rpc/system/internal/global"

type Dictionary struct {
	global.StModel
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`   // 字典名（中）
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`   // 字典名（英）
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`       // 描述
	DictionaryDetails []DictionaryDetail `json:"dictionaryDetails" form:"dictionaryDetails"`
}

func (Dictionary) TableName() string {
	return "st_dictionary"
}

