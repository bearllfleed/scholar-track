package userlogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserListLogic) QueryUserList(in *system.QueryUserListReq) (*system.QueryUserListResp, error) {
	// todo: add your logic here and delete this line

	return &system.QueryUserListResp{}, nil
}
