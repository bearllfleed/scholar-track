package propertyservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPropertyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPropertyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPropertyDetailLogic {
	return &QueryPropertyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryPropertyDetailLogic) QueryPropertyDetail(in *achieve.QueryPropertyDetailReq) (*achieve.QueryPropertyDetailResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.QueryPropertyDetailResp{}, nil
}
