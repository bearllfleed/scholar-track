package property

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPropertyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPropertyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPropertyDetailLogic {
	return &QueryPropertyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPropertyDetailLogic) QueryPropertyDetail(req *types.QueryPropertyDetailReq) (resp *types.QueryPropertyDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
