// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: user.proto

package server

import (
	"context"

	"bookstore/rpc/user/internal/logic/adminuserservice"
	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"
)

type AdminUserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedAdminUserServiceServer
}

func NewAdminUserServiceServer(svcCtx *svc.ServiceContext) *AdminUserServiceServer {
	return &AdminUserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminUserServiceServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	l := adminuserservicelogic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *AdminUserServiceServer) UserList(ctx context.Context, in *user.UserListReq) (*user.UserListResp, error) {
	l := adminuserservicelogic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}
