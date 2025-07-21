package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TResourceModel = (*customTResourceModel)(nil)

type (
	// TResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTResourceModel.
	TResourceModel interface {
		tResourceModel
		withSession(session sqlx.Session) TResourceModel
		FindAll(ctx context.Context) ([]TResource, error)
	}

	customTResourceModel struct {
		*defaultTResourceModel
	}
)

// NewTResourceModel returns a model for the database table.
func NewTResourceModel(conn sqlx.SqlConn) TResourceModel {
	return &customTResourceModel{
		defaultTResourceModel: newTResourceModel(conn),
	}
}

func (m *customTResourceModel) withSession(session sqlx.Session) TResourceModel {
	return NewTResourceModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTResourceModel) FindAll(ctx context.Context) ([]TResource, error) {
	query := fmt.Sprintf("select %s from %s", tResourceRows, m.table)
	var resp []TResource
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return make([]TResource, 0), nil
	default:
		return nil, err
	}
}
