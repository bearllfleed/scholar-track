package initialize

import (
	"github.com/bearllfleed/scholar-track/rpc/storage/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.File{}, &model.FailedDeletion{})
}
func MustNewGrom(dataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	AutoMigrate(db)
	return db
}
