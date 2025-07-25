// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.4

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tAdminUserRoleFieldNames          = builder.RawFieldNames(&TAdminUserRole{}, true)
	tAdminUserRoleRows                = strings.Join(tAdminUserRoleFieldNames, ",")
	tAdminUserRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(tAdminUserRoleFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	tAdminUserRoleRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(tAdminUserRoleFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))
)

type (
	tAdminUserRoleModel interface {
		Insert(ctx context.Context, data *TAdminUserRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TAdminUserRole, error)
		FindOneByUserIdRoleId(ctx context.Context, userId int64, roleId int64) (*TAdminUserRole, error)
		Update(ctx context.Context, data *TAdminUserRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTAdminUserRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TAdminUserRole struct {
		Id        int64     `db:"id"`         // 用户角色ID
		UserId    int64     `db:"user_id"`    // 用户ID
		RoleId    int64     `db:"role_id"`    // 角色ID
		CreatedAt time.Time `db:"created_at"` // 创建时间
		UpdatedAt time.Time `db:"updated_at"` // 更新时间
	}
)

func newTAdminUserRoleModel(conn sqlx.SqlConn) *defaultTAdminUserRoleModel {
	return &defaultTAdminUserRoleModel{
		conn:  conn,
		table: `"public"."t_admin_user_role"`,
	}
}

func (m *defaultTAdminUserRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTAdminUserRoleModel) FindOne(ctx context.Context, id int64) (*TAdminUserRole, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", tAdminUserRoleRows, m.table)
	var resp TAdminUserRole
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTAdminUserRoleModel) FindOneByUserIdRoleId(ctx context.Context, userId int64, roleId int64) (*TAdminUserRole, error) {
	var resp TAdminUserRole
	query := fmt.Sprintf("select %s from %s where user_id = $1 and role_id = $2 limit 1", tAdminUserRoleRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, roleId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTAdminUserRoleModel) Insert(ctx context.Context, data *TAdminUserRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2)", m.table, tAdminUserRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.RoleId)
	return ret, err
}

func (m *defaultTAdminUserRoleModel) Update(ctx context.Context, newData *TAdminUserRole) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, tAdminUserRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Id, newData.UserId, newData.RoleId)
	return err
}

func (m *defaultTAdminUserRoleModel) tableName() string {
	return m.table
}
