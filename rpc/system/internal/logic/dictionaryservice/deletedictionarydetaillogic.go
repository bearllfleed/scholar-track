package dictionaryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(in *system.DeleteDictionaryDetailReq) (*system.DeleteDictionaryDetailResp, error) {
	err := l.svcCtx.DictionaryDetailService.DeleteDictionaryDetail(uint64(in.Id))
	if err != nil {
		return nil, err
	}

	return &system.DeleteDictionaryDetailResp{}, nil
}
