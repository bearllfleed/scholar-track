package casbin

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/client/casbin"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPolicyLogic {
	return &AddPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPolicyLogic) AddPolicy(req *types.AddPolicyReq) error {

	var rpcReqRules []*casbin.Rule
	for _, v := range req.Rules {
		rpcReqRules = append(rpcReqRules, &casbin.Rule{
			Rules: v,
		})
	}
	_, err := l.svcCtx.Casbin.AddPolicy(l.ctx, &casbin.AddPolicyReq{
		Rules: rpcReqRules,
	})
	if err != nil {
		return err
	}
	return nil
}
