package model

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPermissionModel = (*customTPermissionModel)(nil)

type (
	// TPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPermissionModel.
	TPermissionModel interface {
		tPermissionModel
		WithSession(session sqlx.Session) TPermissionModel
		FindAll(ctx context.Context) ([]TPermission, error)
		FindParentAll(ctx context.Context) ([]TPermissionParentData, error)
	}

	customTPermissionModel struct {
		*defaultTPermissionModel
	}
)

type (
	TPermissionParentData struct {
		Id           int64     `db:"id,optional"`          // 权限ID
		Code         int       `db:"code"`                 // 资源编码
		Description  string    `db:"description,optional"` // 权限描述
		PDescription string    `db:"p_description"`        // 父级权限描述
		ParentCode   int       `db:"parent_code"`          // 父级资源编码
		CreatedAt    time.Time `db:"created_at,optional"`  // 创建时间
		UpdatedAt    time.Time `db:"updated_at,optional"`  // 更新时间
	}
)

// NewTPermissionModel returns a model for the database table.
func NewTPermissionModel(conn sqlx.SqlConn) TPermissionModel {
	return &customTPermissionModel{
		defaultTPermissionModel: newTPermissionModel(conn),
	}
}

func (m *customTPermissionModel) WithSession(session sqlx.Session) TPermissionModel {
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

func (m *defaultTPermissionModel) FindParentAll(ctx context.Context) ([]TPermissionParentData, error) {
	query := "select  COALESCE(t_permission.id, 0) as id, t_resource.code, COALESCE(t_permission.description, '') as description, t_resource.description as p_description, t_resource.parent_code as parent_code, t_resource.created_at, t_resource.updated_at from t_resource left outer join t_permission on t_permission.resource_name = t_resource.name;"
	var resp []TPermissionParentData
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return make([]TPermissionParentData, 0), nil
	default:
		return nil, err
	}
}
