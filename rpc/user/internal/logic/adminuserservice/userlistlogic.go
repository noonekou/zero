package adminuserservicelogic

import (
	"context"
	"database/sql"

	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *user.UserListReq) (*user.UserListResp, error) {
	total, err := l.svcCtx.AdminUserModel.Count(l.ctx)
	if err != nil {
		return nil, err
	}

	list, err := l.svcCtx.AdminUserModel.FindAllByPage(l.ctx, in.Page, in.PageSize)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if list == nil {
		return &user.UserListResp{List: nil}, nil
	}

	listData := make([]*user.UserInfo, 0)
	for _, v := range *list {
		listData = append(listData, &user.UserInfo{Id: v.Id, UserName: v.Username, NickName: v.Nickname, Avatar: v.Avatar, Email: v.Email, Phone: v.Phone, Status: int32(v.Status), CreatedAt: v.CreatedAt.Unix(), UpdatedAt: v.UpdatedAt.Unix()})
	}

	return &user.UserListResp{List: listData, Total: total}, nil
}
