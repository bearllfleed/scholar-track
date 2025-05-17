package user

import (
	"context"
	"github.com/bearllfleed/scholar-track/rpc/system/client/user"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Realname: req.Realname,
		Major:    req.Major,
		Grade:    req.Grade,
		College:  req.College,
		Class:    req.Class,
		Gender:   req.Gender,
		RoleId:   req.RoleId,
		Password: utils.BcryptHash(req.Password),
	})
	if err != nil {
		return nil, err
	}
	resp = &types.RegisterResp{
		User: &types.QueryUserDetailResp{
			Id:        registerResp.User.Id,
			Username:  registerResp.User.Username,
			Email:     registerResp.User.Email,
			Phone:     registerResp.User.Phone,
			Nickname:  registerResp.User.Nickname,
			Realname:  registerResp.User.Realname,
			Major:     registerResp.User.Major,
			Grade:     registerResp.User.Grade,
			College:   registerResp.User.College,
			Class:     registerResp.User.Class,
			Gender:    registerResp.User.Gender,
			RoleId:    registerResp.User.RoleId,
			Status:    registerResp.User.Status,
			CreatedAt: registerResp.User.CreatedAt,
			UpdatedAt: registerResp.User.UpdatedAt,
		},
	}
	return
}
