package dictionaryservicelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(in *system.UpdateDictionaryDetailReq) (*system.UpdateDictionaryDetailResp, error) {
	dictionaryDetail := &model.DictionaryDetail{
		StModel: global.StModel{
			Id: uint64(in.Id),
		},
		Label:        in.Label,
		Value:        in.Value,
		Extend:       in.Extend,
		Status:       &in.Status,
		Sort:         int(in.Sort),
		DictionaryId: in.DictionaryId,
	}

	updatedDictionaryDetail, err := l.svcCtx.DictionaryDetailService.UpdateDictionaryDetail(dictionaryDetail)
	if err != nil {
		return nil, err
	}

	return &system.UpdateDictionaryDetailResp{
		DictionaryDetail: &system.DictionaryDetail{
			Id:           uint64(updatedDictionaryDetail.Id),
			Label:        updatedDictionaryDetail.Label,
			Value:        updatedDictionaryDetail.Value,
			Extend:       updatedDictionaryDetail.Extend,
			Status:       *updatedDictionaryDetail.Status,
			Sort:         int32(updatedDictionaryDetail.Sort),
			DictionaryId: updatedDictionaryDetail.DictionaryId,
			CreatedAt:    updatedDictionaryDetail.CreatedAt.Unix(),
			UpdatedAt:    updatedDictionaryDetail.UpdatedAt.Unix(),
		},
	}, nil
}
