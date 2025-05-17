package dictionaryservicelogic

import (
	"context"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/global"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(in *system.UpdateDictionaryReq) (*system.UpdateDictionaryResp, error) {
	dictionary := &model.Dictionary{
		StModel: global.StModel{
			Id: uint64(in.Id),
		},
		Name:   in.Name,
		Type:   in.Type,
		Status: &in.Status,
		Desc:   in.Desc,
	}
	updatedDictionary, err := l.svcCtx.DictionaryService.UpdateDictionary(dictionary)
	if err != nil {
		return nil, err
	}

	return &system.UpdateDictionaryResp{
		Dictionary: &system.Dictionary{
			Id:        uint64(updatedDictionary.Id),
			Name:      updatedDictionary.Name,
			Type:      updatedDictionary.Type,
			Status:    *updatedDictionary.Status,
			Desc:      updatedDictionary.Desc,
			CreatedAt: updatedDictionary.CreatedAt.Unix(),
			UpdatedAt: updatedDictionary.UpdatedAt.Unix(),
		},
	}, nil
}
