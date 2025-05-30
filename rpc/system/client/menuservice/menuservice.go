// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: system.proto

package menuservice

import (
	"context"

	"github.com/bearllfleed/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddPolicyReq                    = system.AddPolicyReq
	AddPolicyResp                   = system.AddPolicyResp
	AddRoleReq                      = system.AddRoleReq
	AddRoleResp                     = system.AddRoleResp
	Api                             = system.Api
	ChangePasswordReq               = system.ChangePasswordReq
	ChangePasswordResp              = system.ChangePasswordResp
	CreateApiReq                    = system.CreateApiReq
	CreateApiResp                   = system.CreateApiResp
	CreateDictionaryDetailReq       = system.CreateDictionaryDetailReq
	CreateDictionaryDetailResp      = system.CreateDictionaryDetailResp
	CreateDictionaryReq             = system.CreateDictionaryReq
	CreateDictionaryResp            = system.CreateDictionaryResp
	CreateMenuReq                   = system.CreateMenuReq
	CreateMenuResp                  = system.CreateMenuResp
	DeleteApiReq                    = system.DeleteApiReq
	DeleteApiResp                   = system.DeleteApiResp
	DeleteDictionaryDetailReq       = system.DeleteDictionaryDetailReq
	DeleteDictionaryDetailResp      = system.DeleteDictionaryDetailResp
	DeleteDictionaryReq             = system.DeleteDictionaryReq
	DeleteDictionaryResp            = system.DeleteDictionaryResp
	DeleteMenuReq                   = system.DeleteMenuReq
	DeleteMenuResp                  = system.DeleteMenuResp
	DeleteRoleReq                   = system.DeleteRoleReq
	DeleteRoleResp                  = system.DeleteRoleResp
	DeleteUserReq                   = system.DeleteUserReq
	DeleteUserResp                  = system.DeleteUserResp
	Dictionary                      = system.Dictionary
	DictionaryDetail                = system.DictionaryDetail
	EnforceReq                      = system.EnforceReq
	EnforceResp                     = system.EnforceResp
	FindUserByUsernameReq           = system.FindUserByUsernameReq
	FindUserByUsernameResp          = system.FindUserByUsernameResp
	Menu                            = system.Menu
	Meta                            = system.Meta
	PolicyInfo                      = system.PolicyInfo
	QueryAllApiReq                  = system.QueryAllApiReq
	QueryAllApiResp                 = system.QueryAllApiResp
	QueryAllDictionaryDetailReq     = system.QueryAllDictionaryDetailReq
	QueryAllDictionaryDetailResp    = system.QueryAllDictionaryDetailResp
	QueryAllDictionaryReq           = system.QueryAllDictionaryReq
	QueryAllDictionaryResp          = system.QueryAllDictionaryResp
	QueryAllMenuTreeReq             = system.QueryAllMenuTreeReq
	QueryAllMenuTreeResp            = system.QueryAllMenuTreeResp
	QueryApiDetailReq               = system.QueryApiDetailReq
	QueryApiDetailResp              = system.QueryApiDetailResp
	QueryDictionaryDetailDetailReq  = system.QueryDictionaryDetailDetailReq
	QueryDictionaryDetailDetailResp = system.QueryDictionaryDetailDetailResp
	QueryDictionaryDetailReq        = system.QueryDictionaryDetailReq
	QueryDictionaryDetailResp       = system.QueryDictionaryDetailResp
	QueryDictionaryListReq          = system.QueryDictionaryListReq
	QueryDictionaryListResp         = system.QueryDictionaryListResp
	QueryMenuDetailReq              = system.QueryMenuDetailReq
	QueryMenuDetailResp             = system.QueryMenuDetailResp
	QueryRoleMenuTreeReq            = system.QueryRoleMenuTreeReq
	QueryRoleMenuTreeResp           = system.QueryRoleMenuTreeResp
	QueryRolePoliciesReq            = system.QueryRolePoliciesReq
	QueryRolePoliciesResp           = system.QueryRolePoliciesResp
	QuerySelfInfoReq                = system.QuerySelfInfoReq
	QuerySelfInfoResp               = system.QuerySelfInfoResp
	QueryUserDetailReq              = system.QueryUserDetailReq
	QueryUserDetailResp             = system.QueryUserDetailResp
	QueryUserListReq                = system.QueryUserListReq
	QueryUserListResp               = system.QueryUserListResp
	RegisterReq                     = system.RegisterReq
	RegisterResp                    = system.RegisterResp
	ResetPasswordReq                = system.ResetPasswordReq
	ResetPasswordResp               = system.ResetPasswordResp
	RoleResp                        = system.RoleResp
	RoleTreeListResp                = system.RoleTreeListResp
	RoleTreeReq                     = system.RoleTreeReq
	RoleTreeResp                    = system.RoleTreeResp
	Rule                            = system.Rule
	SetRolePoliciesReq              = system.SetRolePoliciesReq
	SetRolePoliciesResp             = system.SetRolePoliciesResp
	SetSelfInfoReq                  = system.SetSelfInfoReq
	SetSelfInfoResp                 = system.SetSelfInfoResp
	SetUserInfoReq                  = system.SetUserInfoReq
	SetUserInfoResp                 = system.SetUserInfoResp
	SetUserRoleReq                  = system.SetUserRoleReq
	SetUserRoleResp                 = system.SetUserRoleResp
	UpdateApiReq                    = system.UpdateApiReq
	UpdateApiResp                   = system.UpdateApiResp
	UpdateDictionaryDetailReq       = system.UpdateDictionaryDetailReq
	UpdateDictionaryDetailResp      = system.UpdateDictionaryDetailResp
	UpdateDictionaryReq             = system.UpdateDictionaryReq
	UpdateDictionaryResp            = system.UpdateDictionaryResp
	UpdateMenuReq                   = system.UpdateMenuReq
	UpdateMenuResp                  = system.UpdateMenuResp
	UpdateRoleMenuReq               = system.UpdateRoleMenuReq
	UpdateRoleMenuResp              = system.UpdateRoleMenuResp
	UpdateRoleReq                   = system.UpdateRoleReq
	UpdateRoleResp                  = system.UpdateRoleResp

	MenuService interface {
		CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...grpc.CallOption) (*CreateMenuResp, error)
		QueryAllMenuTree(ctx context.Context, in *QueryAllMenuTreeReq, opts ...grpc.CallOption) (*QueryAllMenuTreeResp, error)
		QueryRoleMenuTree(ctx context.Context, in *QueryRoleMenuTreeReq, opts ...grpc.CallOption) (*QueryRoleMenuTreeResp, error)
		UpdateRoleMenu(ctx context.Context, in *UpdateRoleMenuReq, opts ...grpc.CallOption) (*UpdateRoleMenuResp, error)
		DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...grpc.CallOption) (*DeleteMenuResp, error)
		QueryMenuDetail(ctx context.Context, in *QueryMenuDetailReq, opts ...grpc.CallOption) (*QueryMenuDetailResp, error)
		UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuResp, error)
	}

	defaultMenuService struct {
		cli zrpc.Client
	}
)

func NewMenuService(cli zrpc.Client) MenuService {
	return &defaultMenuService{
		cli: cli,
	}
}

func (m *defaultMenuService) CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...grpc.CallOption) (*CreateMenuResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.CreateMenu(ctx, in, opts...)
}

func (m *defaultMenuService) QueryAllMenuTree(ctx context.Context, in *QueryAllMenuTreeReq, opts ...grpc.CallOption) (*QueryAllMenuTreeResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.QueryAllMenuTree(ctx, in, opts...)
}

func (m *defaultMenuService) QueryRoleMenuTree(ctx context.Context, in *QueryRoleMenuTreeReq, opts ...grpc.CallOption) (*QueryRoleMenuTreeResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.QueryRoleMenuTree(ctx, in, opts...)
}

func (m *defaultMenuService) UpdateRoleMenu(ctx context.Context, in *UpdateRoleMenuReq, opts ...grpc.CallOption) (*UpdateRoleMenuResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.UpdateRoleMenu(ctx, in, opts...)
}

func (m *defaultMenuService) DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...grpc.CallOption) (*DeleteMenuResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

func (m *defaultMenuService) QueryMenuDetail(ctx context.Context, in *QueryMenuDetailReq, opts ...grpc.CallOption) (*QueryMenuDetailResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.QueryMenuDetail(ctx, in, opts...)
}

func (m *defaultMenuService) UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuResp, error) {
	client := system.NewMenuServiceClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}
