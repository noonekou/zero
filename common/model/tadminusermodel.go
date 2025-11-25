package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAdminUserModel = (*customTAdminUserModel)(nil)

type (
	// TAdminUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAdminUserModel.
	TAdminUserModel interface {
		tAdminUserModel
		withSession(session sqlx.Session) TAdminUserModel
		FindOneByUsernameAndPassword(ctx context.Context, username, password string) (*TAdminUser, error)
		FindAllByPage(ctx context.Context, page, pageSize int64) (*[]TAdminUser, error)
		Count(ctx context.Context) (int64, error)
	}

	customTAdminUserModel struct {
		*defaultTAdminUserModel
	}
)

// NewTAdminUserModel returns a model for the database table.
func NewTAdminUserModel(conn sqlx.SqlConn) TAdminUserModel {
	return &customTAdminUserModel{
		defaultTAdminUserModel: newTAdminUserModel(conn),
	}
}

func (m *customTAdminUserModel) withSession(session sqlx.Session) TAdminUserModel {
	return NewTAdminUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTAdminUserModel) FindOneByUsernameAndPassword(ctx context.Context, username, password string) (*TAdminUser, error) {
	var resp TAdminUser
	query := fmt.Sprintf("select %s from %s where username = $1 and password = $2 and status = 1 limit 1", tAdminUserRows, m.table)
	logx.Infof("query: %s, username: %s, password: %s", query, username, password)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username, password)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTAdminUserModel) FindAllByPage(ctx context.Context, page, pageSize int64) (*[]TAdminUser, error) {
	query := fmt.Sprintf("select %s from %s limit $1 offset $2", tAdminUserRows, m.table)
	var resp []TAdminUser
	err := m.conn.QueryRowsCtx(ctx, &resp, query, pageSize, (page-1)*pageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTAdminUserModel) Count(ctx context.Context) (int64, error) {
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
