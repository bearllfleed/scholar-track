package userlogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserRoleLogic {
	return &SetUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserRoleLogic) SetUserRole(in *system.SetUserRoleReq) (*system.SetUserRoleResp, error) {
	// 开启事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 检查角色是否存在
		var count int64
		if err := tx.Model(&model.Role{}).Where("id IN (?)", in.RoleIds).Count(&count).Error; err != nil {
			return err
		}
		if count != int64(len(in.RoleIds)) {
			return cerror.ErrInvalidRole
		}
		// 2. 检查用户是否存在
		if err := tx.Model(&model.User{}).Where("id = ?", in.UserId).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			return cerror.ErrUserNotFound
		}
		// 3. 删除用户角色
		if err := tx.Where("user_id = ?", in.UserId).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		// 4. 添加用户角色
		for _, roleId := range in.RoleIds {
			if err := tx.Create(&model.UserRole{
				UserId: in.UserId,
				RoleId: roleId,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &system.SetUserRoleResp{}, nil
}
