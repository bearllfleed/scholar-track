package model

import (
	"github.com/bearllflee/scholar-track/pkg/global"
)

type File struct {
	global.StModel
	FileKey       string `json:"fileKey" form:"fileKey" gorm:"type:varchar(255);column:file_key;comment:文件key"`
	FileName      string `json:"fileName" form:"fileName" gorm:"type:varchar(255);column:file_name;comment:文件名称"`
	FileType      string `json:"fileType" form:"fileType" gorm:"type:varchar(100);column:file_type;comment:文件类型"`
	FileSize      int64  `json:"fileSize" form:"fileSize" gorm:"type:bigint unsigned;column:file_size;comment:文件大小"`
	BusinessId    uint64 `json:"bussinessId" form:"bussinessId" gorm:"type:bigint unsigned;column:bussiness_id;comment:业务ID"`
	BussinessName string `json:"bussinessName" form:"businessName" gorm:"type:varchar(20);column:bussiness_name;comment:业务名称"`
	FileUrl       string `json:"fileUrl" form:"fileUrl" gorm:"-"`
}

func (File) TableName() string {
	return "st_file"
}
