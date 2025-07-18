package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TApiPermissionModel = (*customTApiPermissionModel)(nil)

type (
	// TApiPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTApiPermissionModel.
	TApiPermissionModel interface {
		tApiPermissionModel
		withSession(session sqlx.Session) TApiPermissionModel
	}

	customTApiPermissionModel struct {
		*defaultTApiPermissionModel
	}
)

// NewTApiPermissionModel returns a model for the database table.
func NewTApiPermissionModel(conn sqlx.SqlConn) TApiPermissionModel {
	return &customTApiPermissionModel{
		defaultTApiPermissionModel: newTApiPermissionModel(conn),
	}
}

func (m *customTApiPermissionModel) withSession(session sqlx.Session) TApiPermissionModel {
	return NewTApiPermissionModel(sqlx.NewSqlConnFromSession(session))
}
