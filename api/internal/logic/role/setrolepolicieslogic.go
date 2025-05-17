package role

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/role"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRolePoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRolePoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRolePoliciesLogic {
	return &SetRolePoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRolePoliciesLogic) SetRolePolicies(req *types.SetRolePoliciesReq) (resp *types.SetRolePoliciesResp, err error) {
	rules := make([]*role.PolicyInfo, 0)
	for _, rule := range req.Rules {
		rules = append(rules, &role.PolicyInfo{
			Path:   rule.Path,
			Method: rule.Method,
		})
	}
	_, err = l.svcCtx.Role.SetRolePolicies(l.ctx, &system.SetRolePoliciesReq{
		RoleId: req.RoleId,
		Rules:  rules,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetRolePoliciesResp{}, nil
}
