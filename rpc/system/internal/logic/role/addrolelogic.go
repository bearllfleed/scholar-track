package rolelogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"strconv"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *system.AddRoleReq) (*system.AddRoleResp, error) {
	var c int64
	var roleModel model.Role
	err := global.DB.Model(&roleModel).Where("role_name = ?", in.RoleName).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, cerror.ErrRoleHasExists
	}
	// 看看父角色是否存在
	if in.ParentId != 0 {
		global.DB.Model(&roleModel).Where("id = ?", in.ParentId).Count(&c)
		if c == 0 {
			return nil, cerror.ErrParentRoleNotExists
		}
	}
	roleModel = model.Role{
		RoleName: in.RoleName,
		ParentId: in.ParentId,
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&roleModel).Error
		if err != nil {
			return err
		}
		// 添加角色权限继承关系
		if in.ParentId != 0 {
			tx.Model(&gormadapter.CasbinRule{}).Create(&gormadapter.CasbinRule{
				Ptype: "g",
				V0:    strconv.Itoa(int(in.ParentId)),
				V1:    strconv.Itoa(int(roleModel.Id)),
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &system.AddRoleResp{
		Role: &system.RoleResp{
			Id:        roleModel.Id,
			RoleName:  roleModel.RoleName,
			ParentId:  roleModel.ParentId,
			CreatedAt: roleModel.CreatedAt.Unix(),
			UpdatedAt: roleModel.UpdatedAt.Unix(),
		},
	}, nil
}
