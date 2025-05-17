package dictionary

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(req *types.DeleteDictionaryReq) (resp *types.DeleteDictionaryResp, err error) {
	_, err = l.svcCtx.DictionaryService.DeleteDictionary(l.ctx, &dictionaryservice.DeleteDictionaryReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteDictionaryResp{}, nil
}
