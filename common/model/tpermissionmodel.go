package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPermissionModel = (*customTPermissionModel)(nil)

type (
	// TPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPermissionModel.
	TPermissionModel interface {
		tPermissionModel
		withSession(session sqlx.Session) TPermissionModel
		FindAll(ctx context.Context) ([]TPermission, error)
	}

	customTPermissionModel struct {
		*defaultTPermissionModel
	}
)

// NewTPermissionModel returns a model for the database table.
func NewTPermissionModel(conn sqlx.SqlConn) TPermissionModel {
	return &customTPermissionModel{
		defaultTPermissionModel: newTPermissionModel(conn),
	}
}

func (m *customTPermissionModel) withSession(session sqlx.Session) TPermissionModel {
	return NewTPermissionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultTPermissionModel) FindAll(ctx context.Context) ([]TPermission, error) {
	query := fmt.Sprintf("select %s from %s", tPermissionRows, m.table)
	var resp []TPermission
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return make([]TPermission, 0), nil
	default:
		return nil, err
	}
}
