package menuservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllMenuTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllMenuTreeLogic {
	return &QueryAllMenuTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllMenuTreeLogic) QueryAllMenuTree(in *system.QueryAllMenuTreeReq) (*system.QueryAllMenuTreeResp, error) {
	menus, err := l.svcCtx.MenuService.GetAllMenuTree(l.ctx)
	if err != nil {
		return nil, err
	}
	var resp system.QueryAllMenuTreeResp
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

func buildMenuTree(menu *model.Menu) []*system.Menu {
	var children []*system.Menu
	for _, child := range menu.Children {
		children = append(children, &system.Menu{
			Id:        child.Id,
			Name:      child.Name,
			Path:      child.Path,
			ParentId:  child.ParentId,
			CreatedAt: child.CreatedAt.Unix(),
			UpdatedAt: child.UpdatedAt.Unix(),
			DeletedAt: child.DeletedAt.Time.Unix(),
			MenuLevel: uint64(child.MenuLevel),
			Hidden:    child.Hidden,
			Component: child.Component,
			Sort:      int32(child.Sort),
			Meta:      &system.Meta{
				Title:      child.Meta.Title,
				Icon:       child.Meta.Icon,
				KeepAlive:  child.Meta.KeepAlive,
				ActiveName: child.Meta.ActiveName,
				CloseTab:   child.Meta.CloseTab,
			},
			Children: buildMenuTree(&child),
		})
	}
	return children

}
