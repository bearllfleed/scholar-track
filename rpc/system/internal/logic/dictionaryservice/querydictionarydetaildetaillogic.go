package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryDetailDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictionaryDetailDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryDetailDetailLogic {
	return &QueryDictionaryDetailDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryDictionaryDetailDetailLogic) QueryDictionaryDetailDetail(in *system.QueryDictionaryDetailDetailReq) (*system.QueryDictionaryDetailDetailResp, error) {
	dictionaryDetail, err := l.svcCtx.DictionaryDetailService.QueryDictionaryDetailDetail(uint64(in.Id))
	if err != nil {
		return nil, err
	}

	return &system.QueryDictionaryDetailDetailResp{
		DictionaryDetail: &system.DictionaryDetail{
			Id:           uint64(dictionaryDetail.Id),
			Label:        dictionaryDetail.Label,
			Value:        dictionaryDetail.Value,
			Extend:       dictionaryDetail.Extend,
			Status:       *dictionaryDetail.Status,
			Sort:         int32(dictionaryDetail.Sort),
			DictionaryId: dictionaryDetail.DictionaryId,
			CreatedAt:    dictionaryDetail.CreatedAt.Unix(),
			UpdatedAt:    dictionaryDetail.UpdatedAt.Unix(),
		},
	}, nil
}
