package service

import (
	"errors"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"gorm.io/gorm"
)

var RoleServiceApp = new(RoleService)

type RoleService struct {
}

func (s *RoleService) QueryRoleById(id uint64) (*model.Role, error) {
	var role model.Role
	err := global.DB.Model(&model.Role{}).Where("id = ?", id).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.ErrRoleNotExists
		}
		return nil, err
	}
	return &role, nil
}
