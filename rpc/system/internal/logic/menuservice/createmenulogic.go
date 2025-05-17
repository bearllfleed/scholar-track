package menuservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMenuLogic) CreateMenu(in *system.CreateMenuReq) (*system.CreateMenuResp, error) {
	total, _, err := l.svcCtx.MenuService.SearchMenu(l.ctx, 1, 10000, in.Name, in.Path, in.Meta.Title, in.Component)
	if err != nil {
		return nil, err
	}
	if total >= 0 {
		return nil, cerror.ErrMenuHasExists
	}
	menu := &model.Menu{
		MenuLevel: 1,
		ParentId:  0,
		Path:      in.Path,
		Name:      in.Name,
		Hidden:    in.Hidden,
		Component: in.Component,
		Sort:      int(in.Sort),
		Meta: model.Meta{
			ActiveName: in.Meta.ActiveName,
			KeepAlive:  in.Meta.KeepAlive,
			Title:      in.Meta.Title,
			Icon:       in.Meta.Icon,
			CloseTab:   in.Meta.CloseTab,
		},
	}
	_, err = l.svcCtx.MenuService.CreateMenu(l.ctx, menu)
	if err != nil {
		return nil, err
	}
	return &system.CreateMenuResp{
		Menu: &system.Menu{
			Id:        menu.Id,
			CreatedAt: menu.CreatedAt.Unix(),
			UpdatedAt: menu.UpdatedAt.Unix(),
			DeletedAt: menu.DeletedAt.Time.Unix(),
			MenuLevel: uint64(menu.MenuLevel),
			ParentId:  uint64(menu.ParentId),
			Path:      menu.Path,
			Name:      menu.Name,
			Hidden:    menu.Hidden,
			Component: menu.Component,
			Sort:      int32(menu.Sort),
			Meta:      &system.Meta{
				ActiveName: menu.Meta.ActiveName,
				KeepAlive:  menu.Meta.KeepAlive,
				Title:      menu.Meta.Title,
				Icon:       menu.Meta.Icon,
				CloseTab:   menu.Meta.CloseTab,
			},
		},
	}, nil
}
