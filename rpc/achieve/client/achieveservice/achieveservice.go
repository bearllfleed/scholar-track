// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: achieve.proto

package achieveservice

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/achieve/achieve"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Category                = achieve.Category
	CreateCategoryReq       = achieve.CreateCategoryReq
	CreateCategoryResp      = achieve.CreateCategoryResp
	CreatePropertyReq       = achieve.CreatePropertyReq
	CreatePropertyResp      = achieve.CreatePropertyResp
	DeleteAchieveReq        = achieve.DeleteAchieveReq
	DeleteAchieveResp       = achieve.DeleteAchieveResp
	DeleteCategoryReq       = achieve.DeleteCategoryReq
	DeleteCategoryResp      = achieve.DeleteCategoryResp
	DeletePropertyReq       = achieve.DeletePropertyReq
	DeletePropertyResp      = achieve.DeletePropertyResp
	Property                = achieve.Property
	QueryCategoryDetailReq  = achieve.QueryCategoryDetailReq
	QueryCategoryListReq    = achieve.QueryCategoryListReq
	QueryCategoryListResp   = achieve.QueryCategoryListResp
	QueryCategoryResp       = achieve.QueryCategoryResp
	QueryPropertyDetailReq  = achieve.QueryPropertyDetailReq
	QueryPropertyDetailResp = achieve.QueryPropertyDetailResp
	QueryPropertyListReq    = achieve.QueryPropertyListReq
	QueryPropertyListResp   = achieve.QueryPropertyListResp
	UpdateCategoryReq       = achieve.UpdateCategoryReq
	UpdateCategoryResp      = achieve.UpdateCategoryResp
	UpdatePropertyReq       = achieve.UpdatePropertyReq
	UpdatePropertyResp      = achieve.UpdatePropertyResp
	UploadAchieveReq        = achieve.UploadAchieveReq
	UploadAchieveResp       = achieve.UploadAchieveResp

	AchieveService interface {
		UploadAchieve(ctx context.Context, in *UploadAchieveReq, opts ...grpc.CallOption) (*UploadAchieveResp, error)
		DeleteAchieve(ctx context.Context, in *DeleteAchieveReq, opts ...grpc.CallOption) (*DeleteAchieveResp, error)
	}

	defaultAchieveService struct {
		cli zrpc.Client
	}
)

func NewAchieveService(cli zrpc.Client) AchieveService {
	return &defaultAchieveService{
		cli: cli,
	}
}

func (m *defaultAchieveService) UploadAchieve(ctx context.Context, in *UploadAchieveReq, opts ...grpc.CallOption) (*UploadAchieveResp, error) {
	client := achieve.NewAchieveServiceClient(m.cli.Conn())
	return client.UploadAchieve(ctx, in, opts...)
}

func (m *defaultAchieveService) DeleteAchieve(ctx context.Context, in *DeleteAchieveReq, opts ...grpc.CallOption) (*DeleteAchieveResp, error) {
	client := achieve.NewAchieveServiceClient(m.cli.Conn())
	return client.DeleteAchieve(ctx, in, opts...)
}
