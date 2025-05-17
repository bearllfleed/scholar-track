package model

import "github.com/bearllflee/scholar-track/rpc/system/internal/global"

type Api struct {
	global.StModel
	Path        string `gorm:"column:path;type:varchar(255);not null;comment:接口路径" json:"path"`
	Method      string `gorm:"column:method;type:varchar(255);not null;comment:请求方法" json:"method"`
	Description string `gorm:"column:description;type:varchar(255);not null;comment:接口描述" json:"description"`
	ApiGroup    string `gorm:"column:api_group;type:varchar(255);not null;comment:接口分组" json:"apiGroup"`
}

func (Api) TableName() string {
	return "st_api"
}
