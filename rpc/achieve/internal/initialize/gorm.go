package initialize

import (
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Category{}, &model.Property{}, &model.AchieveBasic{}, &model.PropertyValue{})
}
func MustNewGorm(dataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	AutoMigrate(db)
	return db
}
