package userlogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByUsernameLogic {
	return &FindUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserByUsernameLogic) FindUserByUsername(in *system.FindUserByUsernameReq) (*system.FindUserByUsernameResp, error) {
	user := &model.User{}
	var c int64
	err := global.DB.Model(user).Where("username = ?", in.Username).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, cerror.ErrUserNotFound
	}
	err = global.DB.Where("username = ?", in.Username).First(user).Error
	if err != nil {
		return nil, err
	}
	return &system.FindUserByUsernameResp{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		RoleId:   user.RoleId,
	}, nil
}
