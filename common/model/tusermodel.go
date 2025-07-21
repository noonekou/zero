package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserModel = (*customTUserModel)(nil)

type (
	// TUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserModel.
	TUserModel interface {
		tUserModel
		withSession(session sqlx.Session) TUserModel
		FindOneByUsernameAndPassword(ctx context.Context, username, password string) (*TUser, error)
		FindAllByPage(ctx context.Context, page, pageSize int64) (*[]TUser, error)
		Count(ctx context.Context) (int64, error)
	}

	customTUserModel struct {
		*defaultTUserModel
	}
)

// NewTUserModel returns a model for the database table.
func NewTUserModel(conn sqlx.SqlConn) TUserModel {
	return &customTUserModel{
		defaultTUserModel: newTUserModel(conn),
	}
}

func (m *customTUserModel) withSession(session sqlx.Session) TUserModel {
	return NewTUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTUserModel) FindOneByUsernameAndPassword(ctx context.Context, username, password string) (*TUser, error) {
	var resp TUser
	query := fmt.Sprintf("select %s from %s where username = $1 and password = $2 limit 1", tUserRows, m.table)
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

func (m *customTUserModel) FindAllByPage(ctx context.Context, page, pageSize int64) (*[]TUser, error) {
	query := fmt.Sprintf("select %s from %s limit $1 offset $2", tUserRows, m.table)
	var resp []TUser
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

func (m *customTUserModel) Count(ctx context.Context) (int64, error) {
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
