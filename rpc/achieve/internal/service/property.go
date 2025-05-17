package service

import (
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/model"
	"gorm.io/gorm"
)

type PropertyService struct {
	db *gorm.DB
}

func NewPropertyService(db *gorm.DB) *PropertyService {
	return &PropertyService{db: db}
}

func CreateProperty(property *model.Property) error {
	return nil
}
