package user

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserDetailLogic) QueryUserDetail(req *types.QueryUserDetailReq) (*types.QueryUserDetailResp, error) {
	resp, err := l.svcCtx.User.QueryUserDetail(l.ctx, &user.QueryUserDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryUserDetailResp{
		Id:        resp.Id,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
		Username:  resp.Username,
		Email:     resp.Email,
		Phone:     resp.Phone,
		Avatar:    resp.Avatar,
		RoleId:    resp.RoleId,
		Status:    resp.Status,
		Nickname:  resp.Nickname,
		Gender:    resp.Gender,
		Major:     resp.Major,
		Grade:     resp.Grade,
		College:   resp.College,
		Realname:  resp.Realname,
		Class:     resp.Class,
		Password:  resp.Password,
	}, nil
}
