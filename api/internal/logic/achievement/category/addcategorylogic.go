package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/achieve/client/categoryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.AddCategoryReq) (resp *types.AddCategoryResp, err error) {
	rpcResp, err := l.svcCtx.Category.CreateCategory(l.ctx, &categoryservice.CreateCategoryReq{
		Name: req.Name,
		Type: req.Type,
		Description: req.Description,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddCategoryResp{
		Id:        rpcResp.Id,
		Name:      rpcResp.Name,
		Type:      rpcResp.Type,
		Description: rpcResp.Description,
		Status: rpcResp.Status,
		CreatedAt: rpcResp.CreatedAt,
		UpdatedAt: rpcResp.UpdatedAt,
		DeletedAt: rpcResp.DeletedAt,
	}, nil
}
