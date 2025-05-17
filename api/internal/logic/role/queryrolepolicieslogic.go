package role

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRolePoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRolePoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRolePoliciesLogic {
	return &QueryRolePoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRolePoliciesLogic) QueryRolePolicies(req *types.QueryRolePoliciesReq) (resp *types.QueryRolePoliciesResp, err error) {
	rolePolicies, err := l.svcCtx.Role.QueryRolePolicies(l.ctx, &system.QueryRolePoliciesReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	policies := make([]*types.PolicyInfo, 0)
	for _, policy := range rolePolicies.Rules {
		policies = append(policies, &types.PolicyInfo{
			Method: policy.Method,
			Path:   policy.Path,
		})
	}

	return &types.QueryRolePoliciesResp{
		Rules: policies,
	}, nil
}
