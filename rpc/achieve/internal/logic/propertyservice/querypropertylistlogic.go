package propertyservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPropertyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPropertyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPropertyListLogic {
	return &QueryPropertyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryPropertyListLogic) QueryPropertyList(in *achieve.QueryPropertyListReq) (*achieve.QueryPropertyListResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.QueryPropertyListResp{}, nil
}
