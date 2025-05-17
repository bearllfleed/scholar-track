package dictionary_detail

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(req *types.UpdateDictionaryDetailReq) (resp *types.UpdateDictionaryDetailResp, err error) {
	_, err = l.svcCtx.DictionaryService.UpdateDictionaryDetail(l.ctx, &dictionaryservice.UpdateDictionaryDetailReq{
		Id:           req.Id,
		DictionaryId: req.DictionaryId,
		Value:        req.Value,
		Sort:         req.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateDictionaryDetailResp{}, nil
}
