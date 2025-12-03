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
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

	err = l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		roleModel := l.svcCtx.RoleModel.WithSession(session)
		rolePermissionModel := l.svcCtx.RolePermissionModel.WithSession(session)
		permissionModel := l.svcCtx.PermissionModel.WithSession(session)

		permissions, err := permissionModel.FindAll(ctx)
		if err != nil {
			return err
		}

		err = rolePermissionModel.DeleteByRoleName(ctx, role.Name)
		if err != nil {
			return err
		}

		toInsertPermission := lo.Filter(permissions, func(v model.TPermission, index int) bool {
			return lo.ContainsBy(in.Permissions, func(v2 *auth.Permission) bool {
				return v.Id == v2.Id
			})
		})

		for _, v := range toInsertPermission {
			_, err = rolePermissionModel.Insert(ctx, &model.TRolePermission{RoleName: in.Name, PermissionName: v.Name})
			if err != nil {
				return err
			}
		}

		err = roleModel.Update(ctx, &model.TRole{Id: in.Id, Name: in.Name, Status: in.Status})

		return err
	})

	return &auth.Empty{}, err
}
