package apiservicelogic

import (
	"context"
	"errors"

	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryApiDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryApiDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryApiDetailLogic {
	return &QueryApiDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryApiDetailLogic) QueryApiDetail(in *system.QueryApiDetailReq) (*system.QueryApiDetailResp, error) {
	api, err := l.svcCtx.ApiService.QueryApiDetail(in.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.ErrApiNotFound
		}
		return nil, err
	}
	return &system.QueryApiDetailResp{
		Id:        api.Id,
		Path:      api.Path,
		Method:    api.Method,
		ApiGroup:  api.ApiGroup,
		Description: api.Description,
		CreatedAt: api.CreatedAt.Unix(),
		UpdatedAt: api.UpdatedAt.Unix(),
	}, nil
}
