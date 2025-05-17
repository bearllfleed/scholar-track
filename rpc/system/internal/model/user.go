package model

import (
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
)

type User struct {
	global.StModel
	Username string `gorm:"type:varchar(15);not null;uniqueIndex:idx_unique_username;comment:学号" json:"username"`
	Email    string `gorm:"type:varchar(255);uniqueIndex:idx_unique_email;comment:邮箱" json:"email"`
	Avatar   string `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	RoleId   uint64 `gorm:"not null;comment:当前角色" json:"roleId"`
	Status   int32  `gorm:"not null;default:1;comment:状态(1:正常, 2:禁用)" json:"status"`
	Nickname string `gorm:"type:varchar(255);comment:昵称" json:"nickname"`
	Phone    string `gorm:"type:char(11);uniqueIndex:idx_unique_phone;comment:手机号" json:"phone"`
	Gender   int32  `gorm:"not null;default:1;comment:性别(1:男, 2:女)" json:"gender"`
	Major    string `gorm:"type:varchar(255);not null;comment:专业" json:"major"`
	College  string `gorm:"type:varchar(255);not null;comment:学院" json:"college"`
	Grade    string `gorm:"type:char(4);not null;comment:年级" json:"grade"`
	Class    string `gorm:"type:varchar(10);not null;comment:班级" json:"class"`
	Realname string `gorm:"type:varchar(255);not null;comment:真实姓名" json:"realname"`
	Password string `gorm:"type:varchar(255);not null;comment:密码" json:"password"`
}

func (User) TableName() string {
	return "st_user"
}
