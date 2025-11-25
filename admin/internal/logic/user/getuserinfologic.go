package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	errs "bookstore/common/error"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	user, err := l.svcCtx.AdminUser.GetUserInfo(l.ctx, &adminuserservice.GetUserInfoReq{Id: l.ctx.Value(types.CtxKeyUserID).(int64)})
	if err != nil {
		return nil, err
	}

	if user.Info == nil {
		return nil, errs.ErrUserNotFound
	}

	// Query user roles
	userRoles, err := l.svcCtx.AdminUserRoleModel.FindAllByUserId(l.ctx, user.Info.Id)
	if err != nil {
		return nil, err
	}

	// Build roles with permissions
	roles := make([]types.RolePermission, 0)
	for _, userRole := range userRoles {
		// Get role information
		role, err := l.svcCtx.RoleModel.FindOne(l.ctx, userRole.RoleId)
		if err != nil {
			continue // Skip if role not found
		}

		// Get role permissions
		permissions, err := l.svcCtx.RolePermissionModel.FindByRoleName(l.ctx, role.Name)
		if err != nil {
			continue // Skip if permissions not found
		}

		// Add each permission as a separate RolePermission entry
		for _, perm := range permissions {
			roles = append(roles, types.RolePermission{
				RoleId:         role.Id,
				RoleName:       role.Name,
				PermissionId:   perm.Id,
				PermissionName: perm.PermissionName,
			})
		}
	}

	return &types.GetUserInfoResp{UserInfo: types.UserInfo{
		Id:        user.Info.Id,
		UserName:  user.Info.UserName,
		NickName:  user.Info.NickName,
		Avatar:    user.Info.Avatar,
		Email:     user.Info.Email,
		Phone:     user.Info.Phone,
		Status:    int(user.Info.Status),
		Roles:     roles,
		CreatedAt: user.Info.CreatedAt,
		UpdatedAt: user.Info.UpdatedAt,
	}}, nil
}
