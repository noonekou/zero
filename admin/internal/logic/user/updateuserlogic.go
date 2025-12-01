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

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) error {

	ids := lo.Map(req.Roles, func(role types.RolePermission, _ int) int64 {
		return role.RoleId
	})

	_, err := l.svcCtx.AdminUser.UpdateUser(l.ctx, &adminuserservice.UserUpdateReq{
		Info: &adminuserservice.UserInfo{
			Id:       req.Id,
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
