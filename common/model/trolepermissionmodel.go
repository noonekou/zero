package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRolePermissionModel = (*customTRolePermissionModel)(nil)

type (
	// TRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRolePermissionModel.
	TRolePermissionModel interface {
		tRolePermissionModel
		withSession(session sqlx.Session) TRolePermissionModel
		FindPermissionNameByUserId(ctx context.Context, userId int64) ([]string, error)
		FindByRoleName(ctx context.Context, roleName string) ([]TPermissionData, error)
		DeleteByRoleName(ctx context.Context, roleName string) error
	}

	customTRolePermissionModel struct {
		*defaultTRolePermissionModel
	}

	TPermissionData struct {
		Id             int64  `db:"id"`              // 权限ID
		PermissionName string `db:"permission_name"` // 权限名
	}
)

// NewTRolePermissionModel returns a model for the database table.
func NewTRolePermissionModel(conn sqlx.SqlConn) TRolePermissionModel {
	return &customTRolePermissionModel{
		defaultTRolePermissionModel: newTRolePermissionModel(conn),
	}
}

func (m *customTRolePermissionModel) withSession(session sqlx.Session) TRolePermissionModel {
	return NewTRolePermissionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultTRolePermissionModel) FindPermissionNameByUserId(ctx context.Context, userId int64) ([]string, error) {
	query := "select permission_name from t_admin_user_role inner join t_role on user_id = $1 and t_role.id = role_id inner join t_role_permission on name = role_name"
	var resp []string
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTRolePermissionModel) FindByRoleName(ctx context.Context, roleName string) ([]TPermissionData, error) {
	query := "select t_permission.id, t_role_permission.permission_name from t_role_permission inner join t_permission on t_permission.name = t_role_permission.permission_name where role_name = $1 order by t_permission.id"
	var resp []TPermissionData
	err := m.conn.QueryRowsCtx(ctx, &resp, query, roleName)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTRolePermissionModel) DeleteByRoleName(ctx context.Context, roleName string) error {
	query := fmt.Sprintf("delete from %s where role_name = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, roleName)
	return err
}
