package model

import (
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
)

type Role struct {
	global.StModel
	RoleName string `gorm:"type:varchar(255);not null;unique" json:"roleName"`
	ParentId uint64 `gorm:"not null;default:0" json:"parentId"`
	Menus    []Menu `json:"menus" gorm:"many2many:st_role_menu;"`
}

func (Role) TableName() string {
	return "st_role"
}
