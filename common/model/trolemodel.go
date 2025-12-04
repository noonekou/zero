package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRoleModel = (*customTRoleModel)(nil)

type (
	// TRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleModel.
	TRoleModel interface {
		tRoleModel
		WithSession(session sqlx.Session) TRoleModel
		FindByPage(ctx context.Context, page, pageSize int64) (*[]TRole, error)
		Count(ctx context.Context) (int64, error)
	}

	customTRoleModel struct {
		*defaultTRoleModel
	}
)

// NewTRoleModel returns a model for the database table.
func NewTRoleModel(conn sqlx.SqlConn) TRoleModel {
	return &customTRoleModel{
		defaultTRoleModel: newTRoleModel(conn),
	}
}

func (m *customTRoleModel) WithSession(session sqlx.Session) TRoleModel {
	return NewTRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTRoleModel) FindByPage(ctx context.Context, page, pageSize int64) (*[]TRole, error) {
	query := fmt.Sprintf("select %s from %s order by created_at desc limit $1 offset $2 ", tRoleRows, m.table)
	var resp []TRole
	err := m.conn.QueryRowsCtx(ctx, &resp, query, pageSize, (page-1)*pageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *customTRoleModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(1) from %s", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
