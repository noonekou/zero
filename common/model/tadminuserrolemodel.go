package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAdminUserRoleModel = (*customTAdminUserRoleModel)(nil)

type (
	// TAdminUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAdminUserRoleModel.
	TAdminUserRoleModel interface {
		tAdminUserRoleModel
		WithSession(session sqlx.Session) TAdminUserRoleModel
		FindAllByUserId(ctx context.Context, userId int64) ([]TAdminUserRole, error)
		DeleteByUserId(ctx context.Context, userId int64) error
		BatchInsert(ctx context.Context, data []*TAdminUserRole) error
		CountByRoleId(ctx context.Context, roleId int64) (int64, error)
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

func (m *customTAdminUserRoleModel) WithSession(session sqlx.Session) TAdminUserRoleModel {
	return NewTAdminUserRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultTAdminUserRoleModel) FindAllByUserId(ctx context.Context, userId int64) ([]TAdminUserRole, error) {
	query := fmt.Sprintf("select %s from %s where user_id = $1", tAdminUserRoleRows, m.table)
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

func (m *defaultTAdminUserRoleModel) DeleteByUserId(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where user_id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

// BatchInsert 批量插入用户角色关系
func (m *defaultTAdminUserRoleModel) BatchInsert(ctx context.Context, data []*TAdminUserRole) error {
	if len(data) == 0 {
		return nil
	}

	// 构建批量插入的 SQL
	// INSERT INTO table (user_id, role_id, status) VALUES ($1, $2, $3), ($4, $5, $6), ...
	valueStrings := make([]string, 0, len(data))
	valueArgs := make([]interface{}, 0, len(data)*3)

	for i, item := range data {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		valueArgs = append(valueArgs, item.UserId, item.RoleId, item.Status)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		m.table,
		tAdminUserRoleRowsExpectAutoSet,
		strings.Join(valueStrings, ","))

	_, err := m.conn.ExecCtx(ctx, query, valueArgs...)
	return err
}

// CountByRoleId 统计指定角色关联的用户数量
func (m *defaultTAdminUserRoleModel) CountByRoleId(ctx context.Context, roleId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE role_id = $1", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, roleId)
	if err != nil {
		return 0, err
	}
	return count, nil
}
