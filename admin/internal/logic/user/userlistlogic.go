package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/common/model"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/samber/lo"
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

		// Query user roles
		rolePermissions, err := l.svcCtx.RolePermissionModel.FindPermissionsByUserId(l.ctx, v.Id)
		if err != nil {
			return nil, err
		}

		rolePermissionsGroup := lo.GroupBy(rolePermissions, func(item model.TRolePermissionData) string {
			return item.RoleName
		})

		roles := make([]types.Role, 0)
		for roleName, rolePermissions := range rolePermissionsGroup {
			permissions := lo.Map(rolePermissions, func(item model.TRolePermissionData, _ int) types.Permission {
				return types.Permission{
					Id:          item.PermissionId,
					Code:        item.PermissionCode, // Converting first char of code string to int
					Description: item.PermissionDescription,
					ParentCode:  int(item.PermissionParentCode),
					Children:    []types.Permission{},
					CreatedAt:   item.CreatedAt.Unix(),
					UpdatedAt:   item.UpdatedAt.Unix(),
				}
			})

			roles = append(roles, types.Role{
				Id:          int64(rolePermissions[0].RoleId),
				Name:        roleName,
				Permissions: permissions,
				CreatedAt:   rolePermissions[0].CreatedAt.Unix(),
				UpdatedAt:   rolePermissions[0].UpdatedAt.Unix(),
			})
		}

		listData = append(listData, types.UserInfo{Id: v.Id, UserName: v.UserName, NickName: v.NickName, Avatar: v.Avatar, Email: v.Email, Phone: v.Phone, Roles: roles, Status: int(v.Status), CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	return &types.UserListResp{List: listData, Total: list.Total}, nil
}
