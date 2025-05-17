package dictionary

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryListLogic {
	return &QueryDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictionaryListLogic) QueryDictionaryList(req *types.QueryDictionaryListReq) (resp *types.QueryDictionaryListResp, err error) {
	rpcResp, err := l.svcCtx.DictionaryService.QueryDictionaryList(l.ctx, &dictionaryservice.QueryDictionaryListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
			Type:     req.Type,
		},
	)
	if err != nil {
		return nil, err
	}

	var dictionaries []*types.Dictionary
	for _, dictionary := range rpcResp.Dictionaries {
		details := make([]*types.DictionaryDetail, len(dictionary.DictionaryDetails))
		for i, detail := range dictionary.DictionaryDetails {
			details[i] = &types.DictionaryDetail{
				Id:           uint64(detail.Id),
				Value:        detail.Value,
				Sort:         int32(detail.Sort),
				DictionaryId: uint64(detail.DictionaryId),
				CreatedAt:    detail.CreatedAt,
				UpdatedAt:    detail.UpdatedAt,
			}
		}
		dictionaries = append(dictionaries, &types.Dictionary{
			Id:        dictionary.Id,
			Name:      dictionary.Name,
			Type:      dictionary.Type,
			Status:    dictionary.Status,
			Desc:      dictionary.Desc,
			DictionaryDetails: details,
			CreatedAt: dictionary.CreatedAt,
			UpdatedAt: dictionary.UpdatedAt,
		})
	}

	return &types.QueryDictionaryListResp{
		Total:    rpcResp.Total,
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
		List: dictionaries,
	}, nil
}
