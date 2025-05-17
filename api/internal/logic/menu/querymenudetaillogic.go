package menu

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/menuservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenuDetailLogic) QueryMenuDetail(req *types.QueryMenuDetailReq) (resp *types.QueryMenuDetailResp, err error) {
	rpcResp, err := l.svcCtx.Menu.QueryMenuDetail(l.ctx, &menuservice.QueryMenuDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryMenuDetailResp{
		Id: rpcResp.Menu.Id,
		Name: rpcResp.Menu.Name,
		Path: rpcResp.Menu.Path,
		ParentId: rpcResp.Menu.ParentId,
		Hidden: rpcResp.Menu.Hidden,
		Component: rpcResp.Menu.Component,
		Sort: int(rpcResp.Menu.Sort),
		Meta: types.Meta{
			Title: rpcResp.Menu.Meta.Title,
			Icon: rpcResp.Menu.Meta.Icon,
			KeepAlive: rpcResp.Menu.Meta.KeepAlive,
			ActiveName: rpcResp.Menu.Meta.ActiveName,
			CloseTab:   rpcResp.Menu.Meta.CloseTab,
		},
		Children: nil,
	}, nil
}
