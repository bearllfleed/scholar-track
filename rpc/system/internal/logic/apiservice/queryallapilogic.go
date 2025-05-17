package apiservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllApiLogic {
	return &QueryAllApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllApiLogic) QueryAllApi(in *system.QueryAllApiReq) (*system.QueryAllApiResp, error) {
	apis, err := l.svcCtx.ApiService.QueryAllApi()
	if err != nil {
		return nil, err
	}
	var apisResp []*system.Api
	for _, api := range apis {
		apisResp = append(apisResp, &system.Api{
			Id:        api.Id,
			Path:      api.Path,
			Method:    api.Method,
			ApiGroup:  api.ApiGroup,
			CreatedAt: api.CreatedAt.Unix(),
		})
	}

	return &system.QueryAllApiResp{
		Apis: apisResp,
	}, nil
}
