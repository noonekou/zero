package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TApiPermissionModel = (*customTApiPermissionModel)(nil)

type (
	// TApiPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTApiPermissionModel.
	TApiPermissionModel interface {
		tApiPermissionModel
		withSession(session sqlx.Session) TApiPermissionModel
		FindOneByMethodAndPath(ctx context.Context, method, path string) (string, error)
	}

	customTApiPermissionModel struct {
		*defaultTApiPermissionModel
	}
)

// NewTApiPermissionModel returns a model for the database table.
func NewTApiPermissionModel(conn sqlx.SqlConn) TApiPermissionModel {
	return &customTApiPermissionModel{
		defaultTApiPermissionModel: newTApiPermissionModel(conn),
	}
}

func (m *customTApiPermissionModel) withSession(session sqlx.Session) TApiPermissionModel {
	return NewTApiPermissionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTApiPermissionModel) FindOneByMethodAndPath(ctx context.Context, method, path string) (string, error) {
	var permissionName string
	query := "select permission_name from t_api_permission inner join t_apis on method = $1 and path = $2 and code = api_code"
	err := m.conn.QueryRowCtx(ctx, &permissionName, query, method, path)
	switch err {
	case nil:
		return permissionName, nil
	case sqlx.ErrNotFound:
		return "", ErrNotFound
	default:
		return "", err
	}
}
