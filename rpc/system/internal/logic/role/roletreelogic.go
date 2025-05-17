package rolelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleTreeLogic {
	return &RoleTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleTreeLogic) RoleTree(in *system.RoleTreeReq) (*system.RoleTreeListResp, error) {
	var topRoles []*model.Role
	err := global.DB.Model(&model.Role{}).Where("parent_id = ?", 0).Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).Find(&topRoles).Error
	if err != nil {
		return nil, err
	}

	// 构建角色树
	var roleTreeList []*system.RoleTreeResp
	for _, topRole := range topRoles {
		roleTree := &system.RoleTreeResp{
			Role: &system.RoleResp{
				Id:       topRole.Id,
				RoleName: topRole.RoleName,
				ParentId: topRole.ParentId,
			},
			Children: []*system.RoleTreeResp{},
		}
		roleTreeList = append(roleTreeList, roleTree)
		l.buildRoleTree(roleTree)
	}

	// 返回结果
	resp := &system.RoleTreeListResp{
		Roles:    roleTreeList,
		Total:    int64(len(roleTreeList)),
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	return resp, nil
}

func (l *RoleTreeLogic) buildRoleTree(parent *system.RoleTreeResp) {
	// 递归构建角色树
	var childRoles []*model.Role
	err := global.DB.Model(&model.Role{}).Where("parent_id = ?", parent.Role.Id).Find(&childRoles).Error
	if err != nil {
		return
	}

	for _, childRole := range childRoles {
		childTree := &system.RoleTreeResp{
			Role: &system.RoleResp{
				Id:       childRole.Id,
				RoleName: childRole.RoleName,
				ParentId: childRole.ParentId,
			},
			Children: []*system.RoleTreeResp{},
		}
		parent.Children = append(parent.Children, childTree)
		l.buildRoleTree(childTree)
	}
}
