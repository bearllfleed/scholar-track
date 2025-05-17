package menuservicelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *system.UpdateMenuReq) (*system.UpdateMenuResp, error) {
	menu, err := l.svcCtx.MenuService.UpdateMenu(l.ctx, &model.Menu{
		StModel: global.StModel{
			Id: in.Id,
		},
		ParentId:  in.ParentId,
		Path:      in.Path,
		Name:      in.Name,
		Hidden:    in.Hidden,
		Component: in.Component,
		Sort:      int(in.Sort),
		Meta: model.Meta{
			Title:      in.Meta.Title,
			Icon:       in.Meta.Icon,
			ActiveName: in.Meta.ActiveName,
			CloseTab:   in.Meta.CloseTab,
			KeepAlive:  in.Meta.KeepAlive,
		},
	})
	if err != nil {
		return nil, err
	}
	return &system.UpdateMenuResp{
		Menu: &system.Menu{
			Id:        menu.Id,
			CreatedAt: menu.CreatedAt.Unix(),
			UpdatedAt: menu.UpdatedAt.Unix(),
			DeletedAt: menu.DeletedAt.Time.Unix(),
			ParentId:  menu.ParentId,
			Path:      menu.Path,
			Name:      menu.Name,
			Hidden:    menu.Hidden,
			Component: menu.Component,
			Sort:      int32(menu.Sort),
			Meta: &system.Meta{
				Title:      menu.Meta.Title,
				Icon:       menu.Meta.Icon,
				ActiveName: menu.Meta.ActiveName,
				CloseTab:   menu.Meta.CloseTab,
				KeepAlive:  menu.Meta.KeepAlive,
			},
		},
	}, nil
}
