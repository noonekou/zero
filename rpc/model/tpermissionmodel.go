package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TPermissionModel = (*customTPermissionModel)(nil)

type (
	// TPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPermissionModel.
	TPermissionModel interface {
		tPermissionModel
		withSession(session sqlx.Session) TPermissionModel
	}

	customTPermissionModel struct {
		*defaultTPermissionModel
	}
)

// NewTPermissionModel returns a model for the database table.
func NewTPermissionModel(conn sqlx.SqlConn) TPermissionModel {
	return &customTPermissionModel{
		defaultTPermissionModel: newTPermissionModel(conn),
	}
}

func (m *customTPermissionModel) withSession(session sqlx.Session) TPermissionModel {
	return NewTPermissionModel(sqlx.NewSqlConnFromSession(session))
}
