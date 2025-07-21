package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TApisModel = (*customTApisModel)(nil)

type (
	// TApisModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTApisModel.
	TApisModel interface {
		tApisModel
		withSession(session sqlx.Session) TApisModel
		FindOneByMethodAndPath(ctx context.Context, method, path string) (*TApis, error)
	}

	customTApisModel struct {
		*defaultTApisModel
	}
)

// NewTApisModel returns a model for the database table.
func NewTApisModel(conn sqlx.SqlConn) TApisModel {
	return &customTApisModel{
		defaultTApisModel: newTApisModel(conn),
	}
}

func (m *customTApisModel) withSession(session sqlx.Session) TApisModel {
	return NewTApisModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultTApisModel) FindOneByMethodAndPath(ctx context.Context, method, path string) (*TApis, error) {
	var resp TApis
	query := fmt.Sprintf("select %s from %s where method = $1 and path = $2 limit 1", tApisRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, method, path)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
