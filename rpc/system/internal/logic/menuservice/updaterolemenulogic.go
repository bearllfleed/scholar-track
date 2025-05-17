package menuservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuLogic {
	return &UpdateRoleMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleMenuLogic) UpdateRoleMenu(in *system.UpdateRoleMenuReq) (*system.UpdateRoleMenuResp, error) {
	err := l.svcCtx.MenuService.UpdateRoleMenu(l.ctx, in.RoleId, in.MenuIds)
	if err != nil {
		return nil, err
	}
	return &system.UpdateRoleMenuResp{}, nil
}