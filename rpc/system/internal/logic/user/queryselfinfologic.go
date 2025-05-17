package userlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySelfInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySelfInfoLogic {
	return &QuerySelfInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySelfInfoLogic) QuerySelfInfo(in *system.QuerySelfInfoReq) (*system.QuerySelfInfoResp, error) {
	var userModel model.User
	err := global.DB.Where(in.Id).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &system.QuerySelfInfoResp{
		User: &system.QueryUserDetailResp{
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
		},
	}, nil

}
