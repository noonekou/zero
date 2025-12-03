package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/samber/lo"
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
	rolePermissions, err := l.svcCtx.RolePermissionModel.FindPermissionsByUserId(l.ctx, user.Info.Id)
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
