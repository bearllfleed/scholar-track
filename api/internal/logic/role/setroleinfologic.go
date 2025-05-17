package role

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/role"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleInfoLogic {
	return &SetRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleInfoLogic) SetRoleInfo(req *types.SetRoleInfoReq) (resp *types.SetRoleInfoResp, err error) {
	_, err = l.svcCtx.Role.UpdateRole(l.ctx, &role.UpdateRoleReq{
		Id:       req.RoleId,
		RoleName: req.RoleName,
	})
	if err != nil {
		return nil, err
	}
	return
}
