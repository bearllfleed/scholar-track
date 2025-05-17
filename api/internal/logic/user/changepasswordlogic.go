package user

import (
	"context"
	"github.com/bearllflee/scholar-track/api/internal/utils"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.ChangePasswordResp, err error) {
	userDetail, err := l.svcCtx.User.QueryUserDetail(l.ctx, &user.QueryUserDetailReq{
		Id: req.UserId,
	})
	if err != nil {
		return
	}
	if !utils.BcryptCheck(req.OldPassword, userDetail.Password) {
		return nil, cerror.ErrPasswordWrong
	}
	_, err = l.svcCtx.User.ChangePassword(l.ctx, &user.ChangePasswordReq{
		UserId:      req.UserId,
		NewPassword: utils.BcryptHash(req.NewPassword),
	})
	return
}
