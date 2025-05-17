package propertyservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropertyLogic {
	return &UpdatePropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePropertyLogic) UpdateProperty(in *achieve.UpdatePropertyReq) (*achieve.UpdatePropertyResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.UpdatePropertyResp{}, nil
}
