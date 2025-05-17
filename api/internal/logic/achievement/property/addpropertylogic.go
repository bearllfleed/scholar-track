package property

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPropertyLogic {
	return &AddPropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPropertyLogic) AddProperty(req *types.AddPropertyReq) (resp *types.AddPropertyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
