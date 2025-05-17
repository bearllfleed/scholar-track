package userlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserDetailLogic) QueryUserDetail(in *system.QueryUserDetailReq) (*system.QueryUserDetailResp, error) {
	var userModel model.User
	var c int64
	err := global.DB.Model(&userModel).Where(in.Id).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, cerror.ErrUserNotFound
	}
	err = global.DB.Where(in.Id).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &system.QueryUserDetailResp{
		Id:        userModel.Id,
		CreatedAt: userModel.CreatedAt.Unix(),
		UpdatedAt: userModel.UpdatedAt.Unix(),
		Username:  userModel.Username,
		Email:     userModel.Email,
		Avatar:    userModel.Avatar,
		RoleId:    userModel.RoleId,
		Status:    userModel.Status,
		Nickname:  userModel.Nickname,
		Phone:     userModel.Phone,
		Gender:    userModel.Gender,
		Major:     userModel.Major,
		College:   userModel.College,
		Grade:     userModel.Grade,
		Class:     userModel.Class,
		Realname:  userModel.Realname,
		Password:  userModel.Password,
	}, nil
}
