package categoryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryListLogic {
	return &QueryCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryCategoryListLogic) QueryCategoryList(in *achieve.QueryCategoryListReq) (*achieve.QueryCategoryListResp, error) {
	err, total, categories := l.svcCtx.CategoryService.QueryCategoryList(l.ctx, in.Name, in.Type, in.Status, int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}
	achieveCategories := make([]*achieve.Category, len(categories))
	for i, category := range categories {
		achieveCategories[i] = &achieve.Category{
			Id:        category.Id,
			Name:      category.Name,
			Type:      category.Type,
			Status:    category.Status,
			CreatedAt: category.CreatedAt.Unix(),
			UpdatedAt: category.UpdatedAt.Unix(),
			DeletedAt: category.DeletedAt.Time.Unix(),
		}
	}
	return &achieve.QueryCategoryListResp{
		Categories: achieveCategories,
		Total:      total,
	}, nil
}
