package propertyservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropertyLogic {
	return &CreatePropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePropertyLogic) CreateProperty(in *achieve.CreatePropertyReq) (*achieve.CreatePropertyResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.CreatePropertyResp{}, nil
}
