package menuservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMenuDetailLogic) QueryMenuDetail(in *system.QueryMenuDetailReq) (*system.QueryMenuDetailResp, error) {
	menu, err := l.svcCtx.MenuService.GetMenuDetail(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &system.QueryMenuDetailResp{
		Menu: &system.Menu{
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
		},
	}, nil
}
