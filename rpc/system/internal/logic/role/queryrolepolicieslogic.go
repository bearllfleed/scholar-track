package rolelogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRolePoliciesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRolePoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRolePoliciesLogic {
	return &QueryRolePoliciesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryRolePoliciesLogic) QueryRolePolicies(in *system.QueryRolePoliciesReq) (*system.QueryRolePoliciesResp, error) {
	// 检查角色是否存在
	var c int64
	err := global.DB.Model(&model.Role{}).Where("id = ?", in.RoleId).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, cerror.ErrRoleNotExists
	}

	var rules []gormadapter.CasbinRule
	err = global.DB.Model(&gormadapter.CasbinRule{}).Where("v0 = ?", in.RoleId).Find(&rules).Error
	if err != nil {
		return nil, err
	}
	policies := make([]*system.PolicyInfo, 0)
	for _, rule := range rules {
		policies = append(policies, &system.PolicyInfo{
			Method: rule.V2,
			Path:   rule.V1,
		})
	}

	return &system.QueryRolePoliciesResp{
		Rules: policies,
	}, nil
}
