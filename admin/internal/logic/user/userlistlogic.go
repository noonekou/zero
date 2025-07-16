package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	list, err := l.svcCtx.AdminUser.UserList(l.ctx, &adminuserservice.UserListReq{Page: req.Page, PageSize: req.PageSize})

	if err != nil {
		return nil, err
	}

	if list == nil {
		return &types.UserListResp{List: nil}, nil
	}

	listData := make([]types.UserInfo, 0)
	for _, v := range list.List {
		listData = append(listData, types.UserInfo{Id: v.Id, UserName: v.UserName, NickName: v.NickName, Avatar: v.Avatar, Email: v.Email, Phone: v.Phone, Status: int(v.Status), CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	return &types.UserListResp{List: listData, Total: list.Total}, nil
}
