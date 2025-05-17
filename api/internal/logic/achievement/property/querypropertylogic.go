package property

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPropertyLogic {
	return &QueryPropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPropertyLogic) QueryProperty(req *types.QueryPropertyReq) (resp *types.QueryPropertyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
