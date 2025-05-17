package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/achieve/client/categoryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryLogic {
	return &QueryCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCategoryLogic) QueryCategory(req *types.QueryCategoryReq) (resp *types.QueryCategoryResp, err error) {
	rpcResp, err := l.svcCtx.Category.QueryCategoryList(l.ctx, &categoryservice.QueryCategoryListReq{
		Page: req.Page,
		PageSize: req.PageSize,
		Name: req.Name,
		Type: req.Type,
		Status: req.Status,
		OrderBy: req.OrderBy,
	})
	if err != nil {
		return nil, err
	}
	var categories []*types.Category
	for _, category := range rpcResp.Categories {
		categories = append(categories, &types.Category{
			Id:        category.Id,
			Name:      category.Name,
			Type:      category.Type,
		})
	}
	return &types.QueryCategoryResp{
		Total: rpcResp.Total,
		Page: rpcResp.Page,
		PageSize: rpcResp.PageSize,
		Categories: categories,
	}, nil
}
