package categoryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *achieve.DeleteCategoryReq) (*achieve.DeleteCategoryResp, error) {
	err := l.svcCtx.CategoryService.DeleteCategory(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &achieve.DeleteCategoryResp{}, nil
}
