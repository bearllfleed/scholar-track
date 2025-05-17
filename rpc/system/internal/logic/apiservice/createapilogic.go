package apiservicelogic

import (
	"context"

	"github.com/bearllfleed/scholar-track/pkg/cerror"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/model"
	"github.com/bearllfleed/scholar-track/rpc/system/internal/svc"
	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateApiLogic) CreateApi(in *system.CreateApiReq) (*system.CreateApiResp, error) {
	// 检查是否存在相同的path 和 method
	exist, err := l.svcCtx.ApiService.CheckApiExist(0, in.Path, in.Method)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, cerror.ErrApiAlreadyExists
	}

	// 创建api
	api := &model.Api{
		Path:   in.Path,
		Method: in.Method,
		Description: in.Description,
		ApiGroup: in.ApiGroup,
	}
	err = l.svcCtx.ApiService.CreateApi(api)
	if err != nil {
		return nil, err
	}
	return &system.CreateApiResp{
		Id:        api.Id,
		CreatedAt: api.CreatedAt.Unix(),
		UpdatedAt: api.UpdatedAt.Unix(),
		DeletedAt: api.DeletedAt.Time.Unix(),
		Path:      api.Path,
		Method:    api.Method,
		ApiGroup:  api.ApiGroup,
		Description: api.Description,
	}, nil
}
