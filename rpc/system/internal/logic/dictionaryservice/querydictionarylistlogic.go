package dictionaryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryListLogic {
	return &QueryDictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryDictionaryListLogic) QueryDictionaryList(in *system.QueryDictionaryListReq) (*system.QueryDictionaryListResp, error) {
	dictionaries, total, err := l.svcCtx.DictionaryService.QueryDictionaryList(in.Page, in.PageSize, in.Name, in.Type)
	if err != nil {
		return nil, err
	}

	resp := &system.QueryDictionaryListResp{
		Total:    total,
		Page:     in.Page,
		PageSize: in.PageSize,
		Dictionaries: make([]*system.Dictionary, 0, len(dictionaries)),
	}

	for _, dictionary := range dictionaries {
		details := make([]*system.DictionaryDetail, len(dictionary.DictionaryDetails))
		for i, detail := range dictionary.DictionaryDetails {
			details[i] = &system.DictionaryDetail{
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
		resp.Dictionaries = append(resp.Dictionaries, &system.Dictionary{
			Id:   uint64(dictionary.Id),
			Name: dictionary.Name,
			Type: dictionary.Type,
			Status: *dictionary.Status,
			Desc: dictionary.Desc,
			DictionaryDetails: details,
			CreatedAt: dictionary.CreatedAt.Unix(),
			UpdatedAt: dictionary.UpdatedAt.Unix(),
		})
	}

	return resp, nil
}
