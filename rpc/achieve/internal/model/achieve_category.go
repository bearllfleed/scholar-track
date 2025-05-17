package model

import (
	"github.com/bearllflee/scholar-track/pkg/global"
)

type Category struct {
	global.StModel
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:名称"`
	Type        string     `json:"type" form:"type" gorm:"column:type;comment:类型(科研、艺术、创新等)"`
	Description string     `json:"description" form:"description" gorm:"column:description;comment:描述"`
	Status      int32      `json:"status" form:"status" gorm:"column:status;comment:状态(1:正常, 2:禁用)"`
	Properties  []Property `json:"properties" form:"properties" gorm:"foreignKey:CategoryId;"`
}

func (Category) TableName() string {
	return "st_category"
}
