// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加用户
func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UserInfo) error {
	ids := lo.Map(req.Roles, func(role types.Role, _ int) int64 {
		return role.Id
	})
	_, err := l.svcCtx.AdminUser.AddUser(l.ctx, &adminuserservice.UserUpdateReq{
		Info: &adminuserservice.UserInfo{
			UserName: req.UserName,
			NickName: req.NickName,
			Avatar:   req.Avatar,
			Email:    req.Email,
			Phone:    req.Phone,
			Status:   int32(req.Status),
		},
		Ids: ids,
	})

	return err
}
