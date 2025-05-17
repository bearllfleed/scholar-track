package menuservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *system.DeleteMenuReq) (*system.DeleteMenuResp, error) {
	err := l.svcCtx.MenuService.DeleteMenu(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &system.DeleteMenuResp{}, nil
}
