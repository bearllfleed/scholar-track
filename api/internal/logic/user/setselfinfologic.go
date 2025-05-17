package user

import (
	"context"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSelfInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSelfInfoLogic {
	return &SetSelfInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSelfInfoLogic) SetSelfInfo(req *types.SetSelfInfoReq) (resp *types.SetSelfInfoResp, err error) {
	_, err = l.svcCtx.User.SetSelfInfo(l.ctx, &user.SetSelfInfoReq{
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
	return &types.SetSelfInfoResp{}, nil
}
