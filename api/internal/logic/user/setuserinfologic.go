package user

import (
	"context"
	"github.com/bearllfleed/scholar-track/rpc/system/client/user"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserInfoLogic) SetUserInfo(req *types.SetUserInfoReq) (resp *types.SetUserInfoResp, err error) {
	_, err = l.svcCtx.User.SetUserInfo(l.ctx, &user.SetUserInfoReq{
		Id:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Major:    req.Major,
		Grade:    req.Grade,
		College:  req.College,
		Realname: req.Realname,
		Class:    req.Class,
		Avatar:   req.Avatar,
		RoleId:   req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetUserInfoResp{}, nil
}
