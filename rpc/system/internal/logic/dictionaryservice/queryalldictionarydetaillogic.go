package dictionaryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllDictionaryDetailLogic {
	return &QueryAllDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllDictionaryDetailLogic) QueryAllDictionaryDetail(in *system.QueryAllDictionaryDetailReq) (*system.QueryAllDictionaryDetailResp, error) {
	dictionaryDetails, err := l.svcCtx.DictionaryDetailService.QueryAllDictionaryDetail()
	if err != nil {
		return nil, err
	}

	resp := &system.QueryAllDictionaryDetailResp{
		DictionaryDetails: make([]*system.DictionaryDetail, len(dictionaryDetails)),
	}

	for i, detail := range dictionaryDetails {
		resp.DictionaryDetails[i] = &system.DictionaryDetail{
			Id:           uint64(detail.Id),
			Label:        detail.Label,
			Value:        detail.Value,
			Extend:       detail.Extend,
			Status:       *detail.Status,
			Sort:         int32(detail.Sort),
			DictionaryId: detail.DictionaryId,
			CreatedAt:    detail.CreatedAt.Unix(),
			UpdatedAt:    detail.UpdatedAt.Unix(),
		}
	}

	return resp, nil
}
