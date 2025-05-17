package user

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySelfInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySelfInfoLogic {
	return &QuerySelfInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySelfInfoLogic) QuerySelfInfo(req *types.QuerySelfInfoReq) (resp *types.QuerySelfInfoResp, err error) {
	info, err := l.svcCtx.User.QuerySelfInfo(l.ctx, &user.QuerySelfInfoReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.QuerySelfInfoResp{
		User: types.QueryUserDetailResp{
			Id:        info.User.Id,
			CreatedAt: info.User.CreatedAt,
			UpdatedAt: info.User.UpdatedAt,
			Username:  info.User.Username,
			Email:     info.User.Email,
			Phone:     info.User.Phone,
			Avatar:    info.User.Avatar,
			RoleId:    info.User.RoleId,
			Status:    info.User.Status,
			Nickname:  info.User.Nickname,
			Gender:    info.User.Gender,
			Major:     info.User.Major,
			Grade:     info.User.Grade,
			College:   info.User.College,
			Realname:  info.User.Realname,
			Class:     info.User.Class,
		},
	}, nil
}
