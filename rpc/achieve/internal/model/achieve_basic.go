package model

import (
	"time"

	"github.com/bearllfleed/scholar-track/pkg/global"
)

// type TeamMembers []uint64

// func (t *TeamMembers) Scan(value interface{}) error {
// 	bytesValue, _ := value.([]byte)
// 	return json.Unmarshal(bytesValue, t)
// }

// func (t TeamMembers) Value() (driver.Value, error) {
// 	return json.Marshal(t)
// }

type AchieveBasic struct {
	global.StModel
	Code        string         `json:"code" form:"code" gorm:"column:code;comment:成果编号;unique"`
	CategoryId  uint64         `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:成果类别"`
	Level       int32          `json:"level" form:"level" gorm:"column:level;comment:获奖级别"`
	Rank        int32          `json:"rank" form:"rank" gorm:"column:rank;comment:获奖排名"`
	AwardTime   time.Time      `json:"awardTime" form:"awardTime" gorm:"column:award_time;comment:获奖时间"`
	Status      int32          `json:"status" form:"status" gorm:"column:status;comment:成果状态(1:保存,2:待审核,3:审核通过,4:审核不通过)"`
	Name        string         `json:"name" form:"name" gorm:"column:name;comment:成果名称"`
	Share       bool           `json:"share" form:"share" gorm:"column:share;comment:是否他人可见"`
	Star        uint64         `json:"star" form:"star" gorm:"column:star;comment:点赞数"`
	Description string         `json:"description" form:"description" gorm:"column:description;comment:成果描述"`
	Uploader    uint64         `json:"uploader" form:"uploader" gorm:"column:uploader;comment:上传者ID"`
	Team        bool           `json:"team" form:"team" gorm:"column:team;comment:是否团队项目"`
	TeamMembers []uint64       `json:"teamMembers" form:"teamMembers" gorm:"column:team_members;comment:团队成员ID;serializer:json"`
	OtherInfo   map[string]any `json:"otherInfo" form:"otherInfo" gorm:"column:other_info;comment:其他信息;type:json;serializer:json"`
}

func (AchieveBasic) TableName() string {
	return "st_achieve_basic"
}
