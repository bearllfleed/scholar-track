package api

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApiLogic) UpdateApi(req *types.UpdateApiReq) (resp *types.UpdateApiResp, err error) {
	rpcResp, err := l.svcCtx.Api.UpdateApi(l.ctx, &system.UpdateApiReq{
		Id: req.Id,
		Path: req.Path,
		Method: req.Method,
		ApiGroup: req.ApiGroup,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UpdateApiResp{
		Id:          rpcResp.Id,
		Path:        rpcResp.Path,
		Method:      rpcResp.Method,
		ApiGroup:    rpcResp.ApiGroup,
		Description: rpcResp.Description,
		CreatedAt:   rpcResp.CreatedAt,
		UpdatedAt:   rpcResp.UpdatedAt,
		DeletedAt:   rpcResp.DeletedAt,
	}

	return resp, nil
}
