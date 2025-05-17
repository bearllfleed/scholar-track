package apiservicelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateApiLogic) UpdateApi(in *system.UpdateApiReq) (*system.UpdateApiResp, error) {
	// 检查api是否存在
	api, err := l.svcCtx.ApiService.QueryApiDetail(in.Id)
	if err != nil {
		return nil, err
	}

	global.DB.Transaction(func(tx *gorm.DB) error {
		// 更新casbin
		err = l.svcCtx.CasbinService.UpdateApi(api.Path, api.Method, in.Path, in.Method)
		if err != nil {
			return err
		}

		api.Path = in.Path
		api.Method = in.Method
		api.ApiGroup = in.ApiGroup
		api.Description = in.Description
		// 更新api
		api, err = l.svcCtx.ApiService.UpdateApi(api)
		if err != nil {
			return err
		}

		return nil
	})

	return &system.UpdateApiResp{
		Id:          api.Id,
		Path:        api.Path,
		Method:      api.Method,
		ApiGroup:    api.ApiGroup,
		Description: api.Description,
		CreatedAt:   api.CreatedAt.Unix(),
		UpdatedAt:   api.UpdatedAt.Unix(),
		DeletedAt:   api.DeletedAt.Time.Unix(),
	}, nil
}
