package file

import (
	"context"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage_client"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq) (resp *types.UploadFileResp, err error) {
	rpcResp, err := l.svcCtx.Storage.FileUpload(l.ctx, &storage_client.FileUploadRequest{
		FileName:      req.FileName,
		FileType:      req.FileType,
		FileData:      req.File,
		BussinessId:   req.BusinessId,
		BussinessName: req.BussinessName,
	})
	if err != nil {
		return nil, err
	}
	return &types.UploadFileResp{
		FileId: rpcResp.FileId,
	}, nil
}
