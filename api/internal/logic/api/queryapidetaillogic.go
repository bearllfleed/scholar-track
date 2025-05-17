package api

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/apiservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryApiDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryApiDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryApiDetailLogic {
	return &QueryApiDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryApiDetailLogic) QueryApiDetail(req *types.QueryApiDetailReq) (resp *types.QueryApiDetailResp, err error) {
	rpcResp, err := l.svcCtx.Api.QueryApiDetail(l.ctx, &apiservice.QueryApiDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.QueryApiDetailResp{
		Id:          rpcResp.Id,
		Path:        rpcResp.Path,
		Method:      rpcResp.Method,
		ApiGroup:    rpcResp.ApiGroup,
		Description: rpcResp.Description,
		CreatedAt:   rpcResp.CreatedAt,
		UpdatedAt:   rpcResp.UpdatedAt,
	}

	return
}
