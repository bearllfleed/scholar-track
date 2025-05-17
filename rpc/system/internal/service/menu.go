package service

import (
	"context"
	"errors"

	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"gorm.io/gorm"
)

var MenuServiceApp = new(MenuService)

type MenuService struct {
}

func (s *MenuService) UpdateRoleMenu(ctx context.Context, roleId uint64, menuIds []uint64) error {
	return global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除与 roleId 相关的所有 RoleMenu 记录
		if err := tx.Where("role_id = ?", roleId).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}

		// 批量创建新的 RoleMenu 记录
		for _, menuId := range menuIds {
			if err := tx.Create(&model.RoleMenu{RoleId: roleId, MenuId: menuId}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *MenuService) CreateMenu(ctx context.Context, menu *model.Menu) (*model.Menu, error) {
	err := global.DB.WithContext(ctx).Create(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s *MenuService) GetMenuDetail(ctx context.Context, id uint64) (*model.Menu, error) {
	var menu model.Menu
	err := global.DB.WithContext(ctx).Where("id = ?", id).First(&menu).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.ErrMenuNotFound
		}
		return nil, err
	}
	return &menu, nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, menu *model.Menu) (*model.Menu, error) {
	err := global.DB.WithContext(ctx).Save(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s *MenuService) DeleteMenu(ctx context.Context, id uint64) error {
	err := global.DB.WithContext(ctx).Delete(&model.Menu{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *MenuService) GetAllMenuTree(ctx context.Context) ([]model.Menu, error) {
	var menus []model.Menu
	err := global.DB.WithContext(ctx).Model(&model.Menu{}).Where("parent_id = ?", 0).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	for i := range menus {
		children, err := s.GetMenuTree(ctx, menus[i].Id)
		if err != nil {
			return nil, err
		}
		menus[i].Children = children
	}
	return menus, nil
}

func (s *MenuService) GetMenuTree(ctx context.Context, id uint64) ([]model.Menu, error) {
	var menus []model.Menu
	err := global.DB.WithContext(ctx).Model(&model.Menu{}).Where("parent_id = ?", id).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	for i := range menus {
		children, err := s.GetMenuTree(ctx, menus[i].Id)
		if err != nil {
			return nil, err
		}
		menus[i].Children = children
	}
	return menus, nil
}

func (s *MenuService) GetMenuTreeByRoleId(ctx context.Context, roleId uint64) ([]model.Menu, error) {
	var menuIds []uint64
	err := global.DB.Model(&model.RoleMenu{}).Where("role_id = ?", roleId).Pluck("menu_id", &menuIds).Error
	if err != nil {
		return nil, err
	}
	var menus []model.Menu
	err = global.DB.Model(&model.Menu{}).Where("id IN (?)", menuIds).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	for i := range menus {
		if menus[i].ParentId == 0 {
			continue
		}
		children, err := s.GetMenuTree(ctx, menus[i].Id)
		if err != nil {
			return nil, err
		}
		menus[i].Children = children
	}
	return menus, nil
}

func (s *MenuService) SearchMenu(ctx context.Context, page, pageSize int, name, path, title, component string) (int64, []*model.Menu, error) {
	var menus []*model.Menu
	db := global.DB.WithContext(ctx).Model(&model.Menu{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if path != "" {
		db = db.Where("path LIKE ?", "%"+path+"%")
	}
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if component != "" {
		db = db.Where("component LIKE ?", "%"+component+"%")
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&menus).Error
	if err != nil {
		return 0, nil, err
	}

	return total, menus, nil
}
