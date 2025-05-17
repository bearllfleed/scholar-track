package casbinlogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPolicyLogic {
	return &AddPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPolicyLogic) AddPolicy(in *system.AddPolicyReq) (*system.AddPolicyResp, error) {
	var rules [][]string
	for _, rule := range in.Rules {
		rules = append(rules, rule.Rules)
	}
	err := l.svcCtx.CasbinService.AddPolicies(rules)
	if err != nil {
		return nil, err
	}
	return &system.AddPolicyResp{}, nil
}
