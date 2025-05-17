package api

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/apiservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiLogic) CreateApi(req *types.CreateApiReq) (resp *types.CreateApiResp, err error) {
	rpcResp, err := l.svcCtx.Api.CreateApi(l.ctx, &apiservice.CreateApiReq{
		Path:      req.Path,
		Method:    req.Method,
		ApiGroup:  req.ApiGroup,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateApiResp{
		Id:        rpcResp.Id,
		CreatedAt: rpcResp.CreatedAt,
		UpdatedAt: rpcResp.UpdatedAt,
		DeletedAt: rpcResp.DeletedAt,
		Path:      rpcResp.Path,
		Method:    rpcResp.Method,
		ApiGroup:  rpcResp.ApiGroup,
		Description: rpcResp.Description,
	}, nil
}
