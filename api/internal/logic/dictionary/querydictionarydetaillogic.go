package dictionary

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryDetailLogic {
	return &QueryDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictionaryDetailLogic) QueryDictionaryDetail(req *types.QueryDictionaryDetailReq) (resp *types.QueryDictionaryDetailResp, err error) {
	rpcResp, err := l.svcCtx.DictionaryService.QueryDictionaryDetail(l.ctx, &dictionaryservice.QueryDictionaryDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryDictionaryDetailResp{
		Id:        rpcResp.Dictionary.Id,
		Name:      rpcResp.Dictionary.Name,
		Type:      rpcResp.Dictionary.Type,
		CreatedAt: rpcResp.Dictionary.CreatedAt,
		UpdatedAt: rpcResp.Dictionary.UpdatedAt,
	}, nil
}
