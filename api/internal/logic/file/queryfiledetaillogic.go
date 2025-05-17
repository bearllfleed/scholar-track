package file

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFileDetailLogic {
	return &QueryFileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryFileDetailLogic) QueryFileDetail(req *types.QueryFileDetailReq) (resp *types.QueryFileDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
