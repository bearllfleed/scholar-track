package initialize

import (
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Api{}, &model.Dictionary{}, &model.DictionaryDetail{}, &model.Menu{}, &model.Role{}, &model.RoleMenu{}, &model.User{}, &model.UserRole{})
}

func MustNewGorm(dataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	AutoMigrate(db)
	return db
}