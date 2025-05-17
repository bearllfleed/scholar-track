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

type SetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserInfoLogic) SetUserInfo(in *system.SetUserInfoReq) (*system.SetUserInfoResp, error) {
	var modelUser model.User
	var c int64
	err := global.DB.Model(&modelUser).Select("id").Where("id = ?", in.Id).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, cerror.ErrUserNotFound
	}
	modelUser = model.User{
		StModel:  global.StModel{Id: in.Id},
		Username: in.Username,
		Email:    in.Email,
		Avatar:   in.Avatar,
		RoleId:   in.RoleId,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Gender:   in.Gender,
		Major:    in.Major,
		College:  in.College,
		Grade:    in.Grade,
		Class:    in.Class,
		Realname: in.Realname,
	}
	err = IfUniqueHasExists(&modelUser)
	if err != nil {
		return nil, err
	}
	err = global.DB.Model(&model.User{}).Select(
		"username", "email", "avatar", "role", "nickname", "phone", "gender", "major", "college", "grade", "class", "realname",
	).Where("id = ?", modelUser.Id).Updates(&modelUser).Error
	if err != nil {
		return nil, err
	}
	return &system.SetUserInfoResp{}, nil
}
