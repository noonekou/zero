package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAdminUserRoleModel = (*customTAdminUserRoleModel)(nil)

type (
	// TAdminUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAdminUserRoleModel.
	TAdminUserRoleModel interface {
		tAdminUserRoleModel
		withSession(session sqlx.Session) TAdminUserRoleModel
		FindAllByUserId(ctx context.Context, userId int64) ([]TAdminUserRole, error)
	}

	customTAdminUserRoleModel struct {
		*defaultTAdminUserRoleModel
	}
)

// NewTAdminUserRoleModel returns a model for the database table.
func NewTAdminUserRoleModel(conn sqlx.SqlConn) TAdminUserRoleModel {
	return &customTAdminUserRoleModel{
		defaultTAdminUserRoleModel: newTAdminUserRoleModel(conn),
	}
}

func (m *customTAdminUserRoleModel) withSession(session sqlx.Session) TAdminUserRoleModel {
	return NewTAdminUserRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultTAdminUserRoleModel) FindAllByUserId(ctx context.Context, userId int64) ([]TAdminUserRole, error) {
	query := fmt.Sprintf("select %s from %s where user_id = $1 and status = 1", tAdminUserRoleRows, m.table)
	var resp []TAdminUserRole
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
