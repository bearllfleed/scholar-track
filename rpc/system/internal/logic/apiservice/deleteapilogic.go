package apiservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteApiLogic) DeleteApi(in *system.DeleteApiReq) (*system.DeleteApiResp, error) {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		api, err := l.svcCtx.ApiService.QueryApiDetail(in.Id)
		if err != nil {
			return err
		}
		l.svcCtx.CasbinService.ClearCasbin(1, api.Path, api.Method)
		err = l.svcCtx.ApiService.DeleteApi(in.Id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &system.DeleteApiResp{}, nil
}
