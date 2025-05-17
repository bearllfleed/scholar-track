package dictionaryservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryLogic {
	return &CreateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryLogic) CreateDictionary(in *system.CreateDictionaryReq) (*system.CreateDictionaryResp, error) {
	dictionary := &model.Dictionary{
		Name:   in.Name,
		Type:   in.Type,
		Status: &in.Status,
		Desc:   in.Desc,
	}
	err := l.svcCtx.DictionaryService.CreateDictionary(dictionary)
	if err != nil {
		return nil, err
	}


	return &system.CreateDictionaryResp{
		Dictionary: &system.Dictionary{
			Id:   uint64(dictionary.Id),
			Name: dictionary.Name,
			Type: dictionary.Type,
			Status: *dictionary.Status,
			Desc: dictionary.Desc,
			CreatedAt: dictionary.CreatedAt.Unix(),
			UpdatedAt: dictionary.UpdatedAt.Unix(),
		},
	}, nil
}
