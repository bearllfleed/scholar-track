package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllDictionaryLogic {
	return &QueryAllDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllDictionaryLogic) QueryAllDictionary(in *system.QueryAllDictionaryReq) (*system.QueryAllDictionaryResp, error) {
	dictionaries, err := l.svcCtx.DictionaryService.QueryAllDictionary()
	if err != nil {
		return nil, err
	}

	resp := &system.QueryAllDictionaryResp{
		Dictionaries: make([]*system.Dictionary, 0, len(dictionaries)),
	}
	for _, dictionary := range dictionaries {
		resp.Dictionaries = append(resp.Dictionaries, &system.Dictionary{
			Id:   uint64(dictionary.Id),
			Name: dictionary.Name,
			Type: dictionary.Type,
			Status: *dictionary.Status,
			Desc: dictionary.Desc,
			CreatedAt: dictionary.CreatedAt.Unix(),
			UpdatedAt: dictionary.UpdatedAt.Unix(),
		})
	}

	return resp, nil
}
