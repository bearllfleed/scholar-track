package api

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/apiservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllApiLogic {
	return &QueryAllApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllApiLogic) QueryAllApi(req *types.QueryAllApiReq) (resp *types.QueryAllApiResp, err error) {
	rpcResp, err := l.svcCtx.Api.QueryAllApi(l.ctx, &apiservice.QueryAllApiReq{})
	if err != nil {
		return nil, err
	}
	var apis []types.Api
	for _, api := range rpcResp.Apis {
		apis = append(apis, types.Api{
			Id:        api.Id,
			Path:      api.Path,
			Method:    api.Method,
			ApiGroup:  api.ApiGroup,
			CreatedAt: api.CreatedAt,
		})
	}
	resp = &types.QueryAllApiResp{
		Apis: apis,
	}
	return
}
