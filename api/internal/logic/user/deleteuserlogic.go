package user

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	_, err = l.svcCtx.User.QueryUserDetail(l.ctx, &user.QueryUserDetailReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.User.DeleteUser(l.ctx, &user.DeleteUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return
}
