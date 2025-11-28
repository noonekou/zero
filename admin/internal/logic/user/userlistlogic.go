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

func (l *UserListLogic) UserList(req *types.PageReq) (resp *types.UserListResp, err error) {
	list, err := l.svcCtx.AdminUser.UserList(l.ctx, &adminuserservice.UserListReq{Page: req.Page, PageSize: req.PageSize})

	if err != nil {
		return nil, err
	}

	if list == nil {
		return &types.UserListResp{List: nil}, nil
	}

	listData := make([]types.UserInfo, 0)
	for _, v := range list.List {

		userRoles, err := l.svcCtx.AdminUserRoleModel.FindAllByUserId(l.ctx, v.Id)
		if err != nil {
			return nil, err
		}

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

		listData = append(listData, types.UserInfo{Id: v.Id, UserName: v.UserName, NickName: v.NickName, Avatar: v.Avatar, Email: v.Email, Phone: v.Phone, Roles: roles, Status: int(v.Status), CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	return &types.UserListResp{List: listData, Total: list.Total}, nil
}
