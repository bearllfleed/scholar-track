package menu

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/menuservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuLogic {
	return &UpdateRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleMenuLogic) UpdateRoleMenu(req *types.UpdateRoleMenuReq) (resp *types.UpdateRoleMenuResp, err error) {
	_, err = l.svcCtx.Menu.UpdateRoleMenu(l.ctx, &menuservice.UpdateRoleMenuReq{
		RoleId: req.RoleId,
		MenuIds: req.MenuIds,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateRoleMenuResp{
	}, nil
}
