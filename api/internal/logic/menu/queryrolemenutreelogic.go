package menu

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/menuservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuTreeLogic {
	return &QueryRoleMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRoleMenuTreeLogic) QueryRoleMenuTree(req *types.QueryRoleMenuTreeReq) (resp *types.QueryRoleMenuTreeResp, err error) {
	rpcResp, err := l.svcCtx.Menu.QueryRoleMenuTree(l.ctx, &menuservice.QueryRoleMenuTreeReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	var menus []types.Menu
	for _, menu := range rpcResp.Menus {
		menus = append(menus, types.Menu{
			Id:        menu.Id,
			Name:      menu.Name,
			Path:      menu.Path,
			ParentId:  menu.ParentId,
			Hidden:    menu.Hidden,
			Component: menu.Component,
			Sort:      int(menu.Sort),
			Meta: types.Meta{
				Title:      menu.Meta.Title,
				Icon:       menu.Meta.Icon,
				KeepAlive:  menu.Meta.KeepAlive,
				ActiveName: menu.Meta.ActiveName,
				CloseTab:   menu.Meta.CloseTab,
			},
			Children: buildMenuTree(menu),
		})
	}
	return &types.QueryRoleMenuTreeResp{
		Menus: menus,
	}, nil
}
