package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryDetailLogic {
	return &QueryDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryDictionaryDetailLogic) QueryDictionaryDetail(in *system.QueryDictionaryDetailReq) (*system.QueryDictionaryDetailResp, error) {
	dictionary, err := l.svcCtx.DictionaryService.QueryDictionaryDetail(uint64(in.Id))
	if err != nil {
		return nil, err
	}

	dictionaryDetails := make([]*system.DictionaryDetail, len(dictionary.DictionaryDetails))
	for i, detail := range dictionary.DictionaryDetails {
		dictionaryDetails[i] = &system.DictionaryDetail{
			Id:           uint64(detail.Id),
			Label:        detail.Label,
			Value:        detail.Value,
			Extend:       detail.Extend,
			Status:       *detail.Status,
			Sort:         int32(detail.Sort),
			DictionaryId: uint64(detail.DictionaryId),
			CreatedAt:    detail.CreatedAt.Unix(),
			UpdatedAt:    detail.UpdatedAt.Unix(),
		}
	}

	return &system.QueryDictionaryDetailResp{
		Dictionary: &system.Dictionary{
			Id:                uint64(dictionary.Id),
			Name:              dictionary.Name,
			Type:              dictionary.Type,
			Status:            *dictionary.Status,
			Desc:              dictionary.Desc,
			DictionaryDetails: dictionaryDetails,
			CreatedAt:         dictionary.CreatedAt.Unix(),
			UpdatedAt:         dictionary.UpdatedAt.Unix(),
		},
	}, nil
}
