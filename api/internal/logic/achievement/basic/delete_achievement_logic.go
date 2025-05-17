package basic

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAchievementLogic struct {
	logx.Logger
	svc *svc.ServiceContext
	ctx context.Context
}

func NewDeleteAchievementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAchievementLogic {
	return &DeleteAchievementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svcCtx,
	}
}

func (l *DeleteAchievementLogic) DeleteAchievement(req *types.DeleteAchievementReq) (resp *types.DeleteAchievementResp, err error) {
	_, err = l.svc.Achieve.DeleteAchieve(l.ctx, &achieve.DeleteAchieveReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
