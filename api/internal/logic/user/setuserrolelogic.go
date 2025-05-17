package user

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserRoleLogic {
	return &SetUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserRoleLogic) SetUserRole(req *types.SetUserRoleReq) (resp *types.SetUserRoleResp, err error) {
	_, err = l.svcCtx.User.SetUserRole(l.ctx, &user.SetUserRoleReq{
		UserId:  req.UserId,
		RoleIds: req.RoleIds,
	})
	return
}
