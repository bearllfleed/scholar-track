package basic

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"
	"github.com/bearllfleed/scholar-track/rpc/storage/storage_client" // Assuming storage client path
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAchievementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadAchievementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAchievementLogic {
	return &UploadAchievementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UploadAchievement now accepts file headers and handles uploading them to storage.
func (l *UploadAchievementLogic) UploadAchievement(req *types.UploadAchievementReq, filesMap map[string][]*multipart.FileHeader) (resp *types.UploadAchievementResp, err error) {
	// uploadedFiles := make([]uint64, 0)
	businessName := "achievement"

	otherInfoPb, err := structpb.NewStruct(req.OtherInfo)
	if err != nil {
		l.Logger.Errorf("Failed to convert OtherInfo to structpb: %v", err)
		return nil, fmt.Errorf("处理 OtherInfo 失败: %v", err)
	}

	achieveReq := &achieve.UploadAchieveReq{
		Code:        req.Code,
		Name:        req.Name,
		Status:      req.Status,
		Level:       req.Level,
		Rank:        req.Rank,
		AwardTime:   req.AwardTime,
		Share:       req.Share,
		Star:        req.Star,
		Uploader:    req.Uploader, // Use ID from context
		Team:        req.Team,
		TeamMembers: req.TeamMembers,
		Description: req.Description,
		CategoryId:  req.CategoryId,
		OtherInfo:   otherInfoPb,
	}

	l.Logger.Infof("Calling achieve RPC UploadAchieve for: %s", req.Name)

	uploadAchieveResp, err := l.svcCtx.Achieve.UploadAchieve(l.ctx, achieveReq) // Use the prepared achieveReq variable
	if err != nil {
		l.Logger.Errorf("Achieve RPC UploadAchieve failed for %s: %v", req.Name, err)
		return nil, fmt.Errorf("创建成果记录失败: %v", err)
	}

	l.Logger.Infof("Achieve RPC UploadAchieve success for: %s, ID: %d", req.Name, uploadAchieveResp.Id)

	for _, fileHeaders := range filesMap {
		for _, header := range fileHeaders {
			file, err := header.Open()
			if err != nil {
				l.Logger.Errorf("Failed to open file header %s: %v", header.Filename, err)
				return nil, fmt.Errorf("打开文件失败: %s (%v)", header.Filename, err)
			}
			defer file.Close() // Ensure file is closed

			fileBytes, err := io.ReadAll(file)
			if err != nil {
				l.Logger.Errorf("Failed to read file content %s: %v", header.Filename, err)
				return nil, fmt.Errorf("读取文件失败: %s (%v)", header.Filename, err)
			}

			uploadReq := &storage_client.FileUploadRequest{
				FileData:      fileBytes,
				FileName:      header.Filename,
				FileType:      header.Header.Get("Content-Type"),
				BussinessId:   uploadAchieveResp.Id,
				BussinessName: businessName,
			}

			l.Logger.Infof("Calling storage RPC FileUpload for: %s, Size: %d, Type: %s", header.Filename, header.Size, uploadReq.FileType)
			storageResp, err := l.svcCtx.Storage.FileUpload(l.ctx, uploadReq)
			if err != nil {
				l.Logger.Errorf("Storage RPC FileUpload failed for %s: %v", header.Filename, err)
				return nil, fmt.Errorf("文件上传到存储服务失败: %s (%v)", header.Filename, err)
			}
			l.Logger.Infof("Storage RPC FileUpload success for: %s, FileID: %d", header.Filename, storageResp.FileId)
			// uploadedFiles = append(uploadedFiles, storageResp.FileId)
		}
	}
	defer func() {
		if err != nil {
			// l.deleteHasUploadedFiles(uploadedFiles)
			if uploadAchieveResp.Id != 0 {
				l.svcCtx.Achieve.DeleteAchieve(l.ctx, &achieve.DeleteAchieveReq{
					Id: uploadAchieveResp.Id,
				})
			}
		}
	}()

	// --- 3. Return response ---
	resp = &types.UploadAchievementResp{
		Id: uploadAchieveResp.Id,
	}
	return resp, nil
}

// func (l *UploadAchievementLogic) deleteHasUploadedFiles(hasUploadedFiles []uint64) {
// 	for _, fileId := range hasUploadedFiles {
// 		l.svcCtx.Storage.FileDelete(l.ctx, &storage_client.FileDeleteRequest{
// 			FileId: fileId,
// 		})
// 	}
// }
