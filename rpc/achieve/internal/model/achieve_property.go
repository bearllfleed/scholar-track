package model

import "github.com/bearllflee/scholar-track/pkg/global"

type Property struct {
	global.StModel
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称"`
	Key          string `json:"key" form:"key" gorm:"column:key;comment:键"`
	Description  string `json:"description" form:"description" gorm:"column:description;comment:描述"`
	Status       int32  `json:"status" form:"status" gorm:"column:status;comment:状态(1:正常, 2:禁用)"`
	Type         string `json:"type" form:"type" gorm:"column:type;comment:属性类型"`
	IsSearch     bool   `json:"isSearch" form:"is_search" gorm:"column:is_search;comment:是否可搜索"`
	IsRequired   bool   `json:"isRequired" form:"is_required" gorm:"column:is_required;comment:是否必填"`
	ValidateRule string `json:"validateRule" form:"validateRule" gorm:"column:validate_rule;comment:验证规则"`
	CategoryId   uint64 `json:"categoryId" form:"category_id" gorm:"column:category_id;comment:所属类别"`
}

func (Property) TableName() string {
	return "st_property"
}
