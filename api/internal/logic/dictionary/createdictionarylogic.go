package dictionary

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryLogic {
	return &CreateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDictionaryLogic) CreateDictionary(req *types.CreateDictionaryReq) (resp *types.CreateDictionaryResp, err error) {
	// 1. 添加字典
	rpcResp, err := l.svcCtx.DictionaryService.CreateDictionary(l.ctx, &dictionaryservice.CreateDictionaryReq{
		Name:  req.Name,
		Type:  req.Type,
		Status: req.Status,
		Desc:  req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateDictionaryResp{
		Id:        rpcResp.Dictionary.Id,
		Name:      rpcResp.Dictionary.Name,
		Type:      rpcResp.Dictionary.Type,
		Status:    rpcResp.Dictionary.Status,
		Desc:      rpcResp.Dictionary.Desc,
		CreatedAt: rpcResp.Dictionary.CreatedAt,
		UpdatedAt: rpcResp.Dictionary.UpdatedAt,
	}, nil
}
