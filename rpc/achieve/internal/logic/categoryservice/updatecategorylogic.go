package categoryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/model"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *achieve.UpdateCategoryReq) (*achieve.UpdateCategoryResp, error) {
	var category model.Category
	category.Id = in.Id
	category.Name = in.Name
	category.Type = in.Type
	category.Status = in.Status
	err := l.svcCtx.CategoryService.UpdateCategory(l.ctx, &category)
	if err != nil {
		return nil, err
	}

	return &achieve.UpdateCategoryResp{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Unix(),
		UpdatedAt: category.UpdatedAt.Unix(),
		Status:    category.Status,
	}, nil
}
