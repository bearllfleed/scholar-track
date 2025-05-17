package role

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/client/role"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.AddRoleReq) (*types.AddRoleResp, error) {
	// 1. 添加角色
	resp, err := l.svcCtx.Role.AddRole(l.ctx, &role.AddRoleReq{
		RoleName: req.RoleName,
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddRoleResp{
		RoleName:  resp.Role.RoleName,
		ParentId:  resp.Role.ParentId,
		CreatedAt: resp.Role.CreatedAt,
		UpdatedAt: resp.Role.UpdatedAt,
		Id:        resp.Role.Id,
	}, nil
}
