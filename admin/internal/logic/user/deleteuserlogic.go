// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.UserInfoReq) error {
	_, err := l.svcCtx.AdminUser.DeleteUser(l.ctx, &adminuserservice.UserInfoReq{
		Id: req.Id,
	})

	if err != nil {
		return err
	}

	err = l.svcCtx.AdminUserRoleModel.DeleteByUserId(l.ctx, req.Id)

	return err
}
