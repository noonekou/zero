package adminauthservicelogic

import (
	"context"
	"database/sql"

	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *auth.Role) (*auth.Empty, error) {
	if in.Id == 0 {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	if in.Name == "" {
		return nil, errs.ErrRoleNameCannotBeEmpty.GRPCStatus().Err()
	}

	if len(in.Permissions) == 0 {
		return nil, errs.ErrPermissionNotFound.GRPCStatus().Err()
	}

	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if role == nil {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	if in.Name != role.Name {
		temp, err := l.svcCtx.RoleModel.FindOneByName(l.ctx, in.Name)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if temp != nil {
			return nil, errs.ErrRoleAlreadyExist.GRPCStatus().Err()
		}
	}

	permissions, err := l.svcCtx.PermissionModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.RolePermissionModel.DeleteByRoleName(l.ctx, role.Name)
	if err != nil {
		return nil, err
	}

	toInsertPermission := lo.Filter(permissions, func(v model.TPermission, index int) bool {
		return lo.ContainsBy(in.Permissions, func(v2 *auth.Permission) bool {
			return v.Id == v2.Id
		})
	})

	for _, v := range toInsertPermission {
		_, err = l.svcCtx.RolePermissionModel.Insert(l.ctx, &model.TRolePermission{RoleName: in.Name, PermissionName: v.Name})
		if err != nil {
			return nil, err
		}
	}

	err = l.svcCtx.RoleModel.Update(l.ctx, &model.TRole{Id: in.Id, Name: in.Name})
	if err != nil {
		return nil, err
	}

	return &auth.Empty{}, nil
}
