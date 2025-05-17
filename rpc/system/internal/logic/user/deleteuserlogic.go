package userlogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *system.DeleteUserReq) (*system.DeleteUserResp, error) {
	err := global.DB.Delete(&model.User{}, in.UserId).Error
	if err != nil {
		return nil, err
	}
	return &system.DeleteUserResp{}, nil
}
