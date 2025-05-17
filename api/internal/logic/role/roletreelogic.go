package role

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleTreeLogic {
	return &RoleTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleTreeLogic) RoleTree(req *types.RoleTreeReq) (resp *types.RoleTreeResp, err error) {
	// 调用 gRPC 服务的 RoleTree 方法
	rpcResp, err := l.svcCtx.Role.RoleTree(l.ctx, &role.RoleTreeReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	// 将 gRPC 返回的角色树数据转换为 types.RoleTree
	var roles []*types.RoleTree
	for _, rpcRole := range rpcResp.Roles {
		roles = append(roles, convertRpcRoleToTypeRole(rpcRole))
	}

	// 构建最终的响应
	resp = &types.RoleTreeResp{
		Roles:    roles,
		Total:    rpcResp.Total,
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
	}

	return resp, nil

}

func convertRpcRoleToTypeRole(rpcRole *role.RoleTreeResp) *types.RoleTree {
	return &types.RoleTree{
		RoleName: rpcRole.Role.RoleName,
		ParentId: rpcRole.Role.ParentId,
		CreatedAt: rpcRole.Role.CreatedAt,
		UpdatedAt: rpcRole.Role.UpdatedAt,
		Id:        rpcRole.Role.Id,
		Children: convertRpcChildrenToTypeChildren(rpcRole.Children),
	}
}
// 递归转换子角色
func convertRpcChildrenToTypeChildren(rpcChildren []*role.RoleTreeResp) []*types.RoleTree {
	var children []*types.RoleTree
	for _, rpcChild := range rpcChildren {
		children = append(children, convertRpcRoleToTypeRole(rpcChild))
	}
	return children
}