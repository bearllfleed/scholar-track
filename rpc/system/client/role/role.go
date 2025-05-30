// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: system.proto

package role

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

	Role interface {
		AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error)
		DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error)
		UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error)
		RoleTree(ctx context.Context, in *RoleTreeReq, opts ...grpc.CallOption) (*RoleTreeListResp, error)
		SetRolePolicies(ctx context.Context, in *SetRolePoliciesReq, opts ...grpc.CallOption) (*SetRolePoliciesResp, error)
		QueryRolePolicies(ctx context.Context, in *QueryRolePoliciesReq, opts ...grpc.CallOption) (*QueryRolePoliciesResp, error)
	}

	defaultRole struct {
		cli zrpc.Client
	}
)

func NewRole(cli zrpc.Client) Role {
	return &defaultRole{
		cli: cli,
	}
}

func (m *defaultRole) AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error) {
	client := system.NewRoleClient(m.cli.Conn())
	return client.AddRole(ctx, in, opts...)
}

func (m *defaultRole) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error) {
	client := system.NewRoleClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultRole) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error) {
	client := system.NewRoleClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultRole) RoleTree(ctx context.Context, in *RoleTreeReq, opts ...grpc.CallOption) (*RoleTreeListResp, error) {
	client := system.NewRoleClient(m.cli.Conn())
	return client.RoleTree(ctx, in, opts...)
}

func (m *defaultRole) SetRolePolicies(ctx context.Context, in *SetRolePoliciesReq, opts ...grpc.CallOption) (*SetRolePoliciesResp, error) {
	client := system.NewRoleClient(m.cli.Conn())
	return client.SetRolePolicies(ctx, in, opts...)
}

func (m *defaultRole) QueryRolePolicies(ctx context.Context, in *QueryRolePoliciesReq, opts ...grpc.CallOption) (*QueryRolePoliciesResp, error) {
	client := system.NewRoleClient(m.cli.Conn())
	return client.QueryRolePolicies(ctx, in, opts...)
}
