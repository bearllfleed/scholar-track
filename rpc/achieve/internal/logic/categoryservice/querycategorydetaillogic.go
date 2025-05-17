package categoryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryDetailLogic {
	return &QueryCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryCategoryDetailLogic) QueryCategoryDetail(in *achieve.QueryCategoryDetailReq) (*achieve.QueryCategoryResp, error) {
	category, err := l.svcCtx.CategoryService.QueryCategoryDetail(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &achieve.QueryCategoryResp{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Unix(),
		UpdatedAt: category.UpdatedAt.Unix(),
		Status:    category.Status,
	}, nil
}
