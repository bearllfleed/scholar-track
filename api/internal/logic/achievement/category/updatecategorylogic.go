package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) (resp *types.UpdateCategoryResp, err error) {
	rpcResp, err := l.svcCtx.Category.UpdateCategory(l.ctx, &achieve.UpdateCategoryReq{
		Id: req.Id,
		Name: req.Name,
		Type: req.Type,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateCategoryResp{
		Id:        rpcResp.Id,
		Name:      rpcResp.Name,
		CreatedAt: rpcResp.CreatedAt,
		UpdatedAt: rpcResp.UpdatedAt,
		Status:    rpcResp.Status,
	}, nil
}
