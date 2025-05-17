package dictionaryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(in *system.CreateDictionaryDetailReq) (*system.CreateDictionaryDetailResp, error) {
	dictionaryDetail := &model.DictionaryDetail{
		Label:        in.Label,
		Value:        in.Value,
		Extend:       in.Extend,
		Status:       &in.Status,
		Sort:         int(in.Sort),
		DictionaryId: in.DictionaryId,
	}

	err := l.svcCtx.DictionaryDetailService.CreateDictionaryDetail(dictionaryDetail)
	if err != nil {
		return nil, err
	}

	return &system.CreateDictionaryDetailResp{
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
