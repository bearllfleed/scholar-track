package property

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropertyLogic {
	return &UpdatePropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePropertyLogic) UpdateProperty(req *types.UpdatePropertyReq) (resp *types.UpdatePropertyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
