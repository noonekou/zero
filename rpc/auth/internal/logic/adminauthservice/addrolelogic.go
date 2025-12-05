package adminauthservicelogic

import (
	"context"

	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *auth.Role) (*auth.Empty, error) {
	if in.Name == "" {
		return nil, errs.ErrRoleNameCannotBeEmpty.GRPCStatus().Err()
	}

	if len(in.Permissions) == 0 {
		return nil, errs.ErrPermissionNotFound.GRPCStatus().Err()
	}

	role, _ := l.svcCtx.RoleModel.FindOneByName(l.ctx, in.Name)
	if role != nil {
		return nil, errs.ErrRoleAlreadyExist.GRPCStatus().Err()
	}

	_, err := l.svcCtx.RoleModel.Insert(l.ctx, &model.TRole{Name: in.Name, Status: in.Status})
	if err != nil {
		return nil, err
	}

	permissions, err := l.svcCtx.PermissionModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	toInsertPermission := lo.Filter(permissions, func(v model.TPermission, index int) bool {
		return lo.ContainsBy(in.Permissions, func(v2 *auth.Permission) bool {
			return v.Id == v2.Id
		})
	})

	// 批量插入角色权限关系
	if len(toInsertPermission) > 0 {
		rolePermissions := make([]*model.TRolePermission, 0, len(toInsertPermission))
		for _, permission := range toInsertPermission {
			rolePermissions = append(rolePermissions, &model.TRolePermission{
				RoleName:       in.Name,
				PermissionName: permission.Name,
			})
		}

		err = l.svcCtx.RolePermissionModel.BatchInsert(l.ctx, rolePermissions)
		if err != nil {
			return nil, err
		}
	}

	return &auth.Empty{}, nil
}
