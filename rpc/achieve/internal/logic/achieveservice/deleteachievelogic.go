package achieveservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/achieve/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAchieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAchieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAchieveLogic {
	return &DeleteAchieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAchieveLogic) DeleteAchieve(in *achieve.DeleteAchieveReq) (*achieve.DeleteAchieveResp, error) {
	// 删除基础信息数据(逻辑删除)
	err := l.svcCtx.AchieveService.DeleteAchieve(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// 删除文件信息(物理删除)
	l.svcCtx.StorageClient.DeleteBusinessFiles(l.ctx, &storage.DeleteBusinessFilesRequest{
		BussnessId:   in.Id,
		BusinessName: "achievement",
	})

	return &achieve.DeleteAchieveResp{}, nil
}
