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
		WithSession(session sqlx.Session) TAdminUserModel
		InsertWithId(ctx context.Context, data *TAdminUser) (int64, error)
		FindOneByUsernameAndPassword(ctx context.Context, username, password string) (*TAdminUser, error)
		FindOneByEmailAndPassword(ctx context.Context, email, password string) (*TAdminUser, error)
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

func (m *customTAdminUserModel) WithSession(session sqlx.Session) TAdminUserModel {
	return NewTAdminUserModel(sqlx.NewSqlConnFromSession(session))
}

// InsertWithId 插入数据并返回ID
func (m *defaultTAdminUserModel) InsertWithId(ctx context.Context, data *TAdminUser) (int64, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		m.table, tAdminUserRowsExpectAutoSet)

	var id int64
	err := m.conn.QueryRowCtx(ctx, &id, query,
		data.Username, data.Password, data.Nickname,
		data.Avatar, data.Email, data.Phone, data.Status)

	return id, err
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

func (m *customTAdminUserModel) FindOneByEmailAndPassword(ctx context.Context, email, password string) (*TAdminUser, error) {
	var resp TAdminUser
	query := fmt.Sprintf("select %s from %s where email = $1 and password = $2 and status = 1 limit 1", tAdminUserRows, m.table)
	logx.Infof("query: %s, email: %s, password: %s", query, email, password)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email, password)
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
	query := fmt.Sprintf("select %s from %s order by created_at desc limit $1 offset $2 ", tAdminUserRows, m.table)
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
