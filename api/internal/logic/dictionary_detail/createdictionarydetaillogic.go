package dictionary_detail

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/dictionaryservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(req *types.CreateDictionaryDetailReq) (resp *types.CreateDictionaryDetailResp, err error) {
	// 1. 添加字典项
	rpcResp, err := l.svcCtx.DictionaryService.CreateDictionaryDetail(l.ctx, &dictionaryservice.CreateDictionaryDetailReq{
		DictionaryId: req.DictionaryId,
		Value:        req.Value,
		Sort:         req.Sort,
		Label:        req.Key,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateDictionaryDetailResp{
		Id:           rpcResp.DictionaryDetail.Id,
		DictionaryId: rpcResp.DictionaryDetail.DictionaryId,
		Value:        rpcResp.DictionaryDetail.Value,
		Sort:         rpcResp.DictionaryDetail.Sort,
		CreatedAt:    rpcResp.DictionaryDetail.CreatedAt,
		UpdatedAt:    rpcResp.DictionaryDetail.UpdatedAt,
	}, nil
}
