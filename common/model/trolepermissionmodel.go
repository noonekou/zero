package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TRolePermissionModel = (*customTRolePermissionModel)(nil)

type (
	// TRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRolePermissionModel.
	TRolePermissionModel interface {
		tRolePermissionModel
		withSession(session sqlx.Session) TRolePermissionModel
	}

	customTRolePermissionModel struct {
		*defaultTRolePermissionModel
	}
)

// NewTRolePermissionModel returns a model for the database table.
func NewTRolePermissionModel(conn sqlx.SqlConn) TRolePermissionModel {
	return &customTRolePermissionModel{
		defaultTRolePermissionModel: newTRolePermissionModel(conn),
	}
}

func (m *customTRolePermissionModel) withSession(session sqlx.Session) TRolePermissionModel {
	return NewTRolePermissionModel(sqlx.NewSqlConnFromSession(session))
}
