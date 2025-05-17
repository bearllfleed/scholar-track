package menuservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleMenuTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuTreeLogic {
	return &QueryRoleMenuTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryRoleMenuTreeLogic) QueryRoleMenuTree(in *system.QueryRoleMenuTreeReq) (*system.QueryRoleMenuTreeResp, error) {
	menus, err := l.svcCtx.MenuService.GetMenuTreeByRoleId(l.ctx, in.RoleId)
	if err != nil {
		return nil, err
	}
	var resp system.QueryRoleMenuTreeResp
	resp.Menus = make([]*system.Menu, 0, len(menus))
	for _, menu := range menus {
		resp.Menus = append(resp.Menus, &system.Menu{
			Id:        menu.Id,
			Name:      menu.Name,
			Path:      menu.Path,
			ParentId:  menu.ParentId,
			CreatedAt: menu.CreatedAt.Unix(),
			UpdatedAt: menu.UpdatedAt.Unix(),
			DeletedAt: menu.DeletedAt.Time.Unix(),
			MenuLevel: uint64(menu.MenuLevel),
			Hidden:    menu.Hidden,
			Component: menu.Component,
			Sort:      int32(menu.Sort),
			Meta:      &system.Meta{
				Title:      menu.Meta.Title,
				Icon:       menu.Meta.Icon,
				KeepAlive:  menu.Meta.KeepAlive,
				ActiveName: menu.Meta.ActiveName,
				CloseTab:   menu.Meta.CloseTab,
			},
			Children: buildMenuTree(&menu),
		})
	}
	return &resp, nil
}
