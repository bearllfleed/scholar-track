package user

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserListLogic) QueryUserList(req *types.QueryUserListReq) (*types.QueryUserListResp, error) {
	resp, err := l.svcCtx.User.QueryUserList(l.ctx, &user.QueryUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		RoleId:   req.RoleId,
		Status:   req.Status,
		Realname: req.Realname,
		Major:    req.Major,
		College:  req.College,
		Grade:    req.Grade,
		Class:    req.Class,
		Gender:   req.Gender,
		OrderBy:  req.OrderBy,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	// 考虑resp.Users为空的情况
	if resp.Total == 0 {
		return &types.QueryUserListResp{
			Total:    0,
			List:     []*types.QueryUserDetailResp{},
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}
	var users []*types.QueryUserDetailResp
	for _, user := range resp.Users {
		users = append(users, &types.QueryUserDetailResp{
			Id:        user.Id,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     user.Phone,
			RoleId:    user.RoleId,
			Status:    user.Status,
			Realname:  user.Realname,
			Major:     user.Major,
			College:   user.College,
			Grade:     user.Grade,
			Class:     user.Class,
			Gender:    user.Gender,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return &types.QueryUserListResp{
		Total:    resp.Total,
		List:     users,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
