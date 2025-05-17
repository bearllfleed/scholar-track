package user

import (
	"context"
	"github.com/bearllfleed/scholar-track/api/internal/utils"
	"github.com/bearllfleed/scholar-track/rpc/system/client/user"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) (resp *types.ResetPasswordResp, err error) {
	_, err = l.svcCtx.User.QueryUserDetail(l.ctx, &user.QueryUserDetailReq{Id: req.UserId})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.User.ResetPassword(l.ctx, &user.ResetPasswordReq{
		UserId:   req.UserId,
		Password: utils.BcryptHash("bearllflee"),
	})
	if err != nil {
		return nil, err
	}
	return &types.ResetPasswordResp{}, nil
}
