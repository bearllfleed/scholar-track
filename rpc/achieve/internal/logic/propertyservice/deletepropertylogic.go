package propertyservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePropertyLogic {
	return &DeletePropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePropertyLogic) DeleteProperty(in *achieve.DeletePropertyReq) (*achieve.DeletePropertyResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.DeletePropertyResp{}, nil
}
