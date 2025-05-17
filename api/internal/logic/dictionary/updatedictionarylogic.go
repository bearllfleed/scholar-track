package dictionary

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(req *types.UpdateDictionaryReq) (resp *types.UpdateDictionaryResp, err error) {
	_, err = l.svcCtx.DictionaryService.UpdateDictionary(l.ctx, &dictionaryservice.UpdateDictionaryReq{
		Id:     req.Id,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateDictionaryResp{}, nil
}
