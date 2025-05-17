package model

import "github.com/bearllflee/scholar-track/pkg/global"

type PropertyValue struct {
	global.StModel
	PropertyId uint64 `json:"propertyId" form:"propertyId" gorm:"column:property_id;comment:属性ID"`
	Value      string `json:"value" form:"value" gorm:"column:value;comment:属性值"`
	BasicId    uint64 `json:"basicId" form:"basicId" gorm:"column:basic_id;comment:成果ID"`
}

func (PropertyValue) TableName() string {
	return "st_property_value"
}
