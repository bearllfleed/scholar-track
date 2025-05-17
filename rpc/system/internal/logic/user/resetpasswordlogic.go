package userlogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *system.ResetPasswordReq) (*system.ResetPasswordResp, error) {
	err := global.DB.Model(&model.User{}).Where("id = ?", in.UserId).Update("password", in.Password).Error
	if err != nil {
		return nil, err
	}
	return &system.ResetPasswordResp{}, nil
}
