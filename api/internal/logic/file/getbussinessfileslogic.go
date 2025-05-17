package file

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/storage/storage_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBussinessFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBussinessFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBussinessFilesLogic {
	return &GetBussinessFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBussinessFilesLogic) GetBussinessFiles(req *types.GetBussinessFilesReq) (resp *types.GetBussinessFilesResp, err error) {
	files, err := l.svcCtx.Storage.GetBussinessFiles(l.ctx, &storage_client.GetBussinessFilesRequest{
		BussinessId:   req.BusinessId,
		BussinessName: req.BussinessName,
	})
	if err != nil {
		return nil, err
	}
	filesResponse := make([]*types.File, 0)
	for _, file := range files.Files {
		filesResponse = append(filesResponse, &types.File{
			Id:         file.FileId,
			FileKey:    file.FileKey,
			FileUrl:    file.FileUrl,
			FileName:   file.FileName,
			FileSize:   int64(file.FileSize),
			FileType:   file.FileType,
			BusinessId: file.BussinessId,
			CreatedAt:  file.CreatedAt,
			UpdatedAt:  file.UpdatedAt,
			DeletedAt:  file.DeletedAt,
		})
	}

	return &types.GetBussinessFilesResp{
		Files: filesResponse,
	}, nil
}
