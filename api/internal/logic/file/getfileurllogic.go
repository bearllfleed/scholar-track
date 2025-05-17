package file

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUrlLogic {
	return &GetFileUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileUrlLogic) GetFileUrl(req *types.GetFileUrlReq) (resp *types.GetFileUrlResp, err error) {
	

	return
}
