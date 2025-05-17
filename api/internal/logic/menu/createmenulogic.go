package menu

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/menuservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.CreateMenuReq) (resp *types.CreateMenuResp, err error) {
	rpcResp, err := l.svcCtx.Menu.CreateMenu(l.ctx, &menuservice.CreateMenuReq{
		Name: req.Name,
		Path: req.Path,
		ParentId:  req.ParentId,
		Hidden:    req.Hidden,
		Component: req.Component,
		Sort:      int32(req.Sort),
		Meta: &menuservice.Meta{
			Title:      req.Meta.Title,
			Icon:       req.Meta.Icon,
			KeepAlive:  req.Meta.KeepAlive,
			ActiveName: req.Meta.ActiveName,
			CloseTab:   req.Meta.CloseTab,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateMenuResp{
		Id: rpcResp.Menu.Id,
	}, nil
}

