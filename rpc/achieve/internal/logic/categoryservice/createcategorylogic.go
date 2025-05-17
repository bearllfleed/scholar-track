package categoryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/model"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *achieve.CreateCategoryReq) (*achieve.CreateCategoryResp, error) {
	// 检查名称是否存在
	_, _, categories := l.svcCtx.CategoryService.QueryCategoryList(l.ctx, in.Name, in.Type, 0, 1, 1)
	if len(categories) > 0 {
		return nil, cerror.ErrCategoryNameAlreadyExists
	}
	category, err := l.svcCtx.CategoryService.CreateCategory(l.ctx, &model.Category{
		Name:        in.Name,
		Type:        in.Type,
		Description: in.Description,
		Status:      in.Status,
	})
	if err != nil {
		return nil, err
	}

	return &achieve.CreateCategoryResp{
		Id:        category.Id,
		Name:      category.Name,
		Type:      category.Type,
		Status:    category.Status,
		CreatedAt: category.CreatedAt.Unix(),
		UpdatedAt: category.UpdatedAt.Unix(),
		DeletedAt: category.DeletedAt.Time.Unix(),
	}, nil
}
