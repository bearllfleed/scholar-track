package rolelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRolePoliciesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRolePoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRolePoliciesLogic {
	return &SetRolePoliciesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRolePoliciesLogic) SetRolePolicies(in *system.SetRolePoliciesReq) (*system.SetRolePoliciesResp, error) {
	// 查看角色是否存在
	_, err := l.svcCtx.RoleService.QueryRoleById(in.RoleId)
	if err != nil {
		return nil, err
	}
	// todo: 查看api是否存在
	for _, rule := range in.Rules {
		_, err := l.svcCtx.ApiService.QueryApiByPathAndMethod(rule.Path, rule.Method)
		if err != nil {
			return nil, err
		}
	}
	// authorityId := strconv.Itoa(int(in.RoleId))
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// err := tx.Model(&gormadapter.CasbinRule{}).Where("v0 = ? AND ptype != ?", authorityId, "g").First(&gormadapter.CasbinRule{}).Error
		// if err != nil {
		// 	return err
		// }
		err = l.svcCtx.CasbinService.UpdateCasbin(in.RoleId, in.Rules)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &system.SetRolePoliciesResp{}, nil
}
