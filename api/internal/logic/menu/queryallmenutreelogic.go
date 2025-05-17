package menu

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/client/menuservice"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllMenuTreeLogic {
	return &QueryAllMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllMenuTreeLogic) QueryAllMenuTree(req *types.QueryAllMenuTreeReq) (resp *types.QueryAllMenuTreeResp, err error) {
	rpcResp, err := l.svcCtx.Menu.QueryAllMenuTree(l.ctx, &menuservice.QueryAllMenuTreeReq{})
	if err != nil {
		return nil, err
	}
	var menus []types.Menu
	for _, menu := range rpcResp.Menus {
		menus = append(menus, types.Menu{
			Id: menu.Id,
			Name: menu.Name,
			Path: menu.Path,
			ParentId: menu.ParentId,
			Hidden: menu.Hidden,
			Component: menu.Component,
			Sort: int(menu.Sort),
			Meta: types.Meta{
				Title: menu.Meta.Title,
				Icon: menu.Meta.Icon,
				KeepAlive: menu.Meta.KeepAlive,
				ActiveName: menu.Meta.ActiveName,
				CloseTab: menu.Meta.CloseTab,
			},
			Children: buildMenuTree(menu),
		})
	}
	resp = &types.QueryAllMenuTreeResp{
		Menus: menus,
	}
	return
}

func buildMenuTree(menu *system.Menu) []types.Menu {
	var children []types.Menu
	for _, child := range menu.Children {
		children = append(children, types.Menu{
			Id: child.Id,
			Name: child.Name,
			Path: child.Path,
			ParentId: child.ParentId,
			Hidden: child.Hidden,
			Component: child.Component,
			Sort: int(child.Sort),
			Meta: types.Meta{
				Title:      child.Meta.Title,
				Icon:       child.Meta.Icon,
				KeepAlive:  child.Meta.KeepAlive,
				ActiveName: child.Meta.ActiveName,
				CloseTab:   child.Meta.CloseTab,
			},
			Children: buildMenuTree(child),
		})
	}
	return children
}
