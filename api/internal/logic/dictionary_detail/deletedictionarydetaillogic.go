package dictionary_detail

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/dictionaryservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(req *types.DeleteDictionaryDetailReq) (resp *types.DeleteDictionaryDetailResp, err error) {
	_, err = l.svcCtx.DictionaryService.DeleteDictionaryDetail(l.ctx, &dictionaryservice.DeleteDictionaryDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteDictionaryDetailResp{}, nil
}
