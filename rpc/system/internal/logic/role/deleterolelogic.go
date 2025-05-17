package rolelogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *system.DeleteRoleReq) (*system.DeleteRoleResp, error) {
	// 查询角色是否存在
	var c int64
	var roleModel model.Role
	err := global.DB.Model(&roleModel).Where("id = ?", in.Id).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, cerror.ErrRoleNotExists
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", in.Id).Delete(&roleModel).Error
		if err != nil {
			return err
		}
		// 删除角色权限继承关系
		err = tx.Model(&gormadapter.CasbinRule{}).Where("v1 = ? or v0 = ?", strconv.Itoa(int(roleModel.Id)), strconv.Itoa(int(roleModel.Id))).Delete(&gormadapter.CasbinRule{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &system.DeleteRoleResp{
		Success: true,
	}, nil
}
