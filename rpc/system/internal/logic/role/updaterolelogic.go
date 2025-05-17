package rolelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *system.UpdateRoleReq) (*system.UpdateRoleResp, error) {
	// 查询角色是否存在
	var roleModel model.Role
	var c int64
	// 先查看角色是否存在
	global.DB.Model(&roleModel).Where("id = ?", in.Id).Count(&c)
	if c == 0 {
		return nil, cerror.ErrRoleNotExists
	}
	// 看看修改后的角色名称是否存在
	global.DB.Model(&roleModel).Where("role_name = ? AND id != ?", in.RoleName, in.Id).Count(&c)
	if c > 0 {
		return nil, cerror.ErrRoleHasExists
	}
	roleModel.RoleName = in.RoleName
	roleModel.Id = in.Id
	roleModel.UpdatedAt = time.Now()
	roleModel.CreatedAt = time.Now()
	// 更新角色信息
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&roleModel).Updates(roleModel).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &system.UpdateRoleResp{
		Id:       roleModel.Id,
		RoleName: roleModel.RoleName,
		ParentId: roleModel.ParentId,
	}, nil
}
