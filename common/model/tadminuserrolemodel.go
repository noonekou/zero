package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TAdminUserRoleModel = (*customTAdminUserRoleModel)(nil)

type (
	// TAdminUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAdminUserRoleModel.
	TAdminUserRoleModel interface {
		tAdminUserRoleModel
		withSession(session sqlx.Session) TAdminUserRoleModel
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

func (m *customTAdminUserRoleModel) withSession(session sqlx.Session) TAdminUserRoleModel {
	return NewTAdminUserRoleModel(sqlx.NewSqlConnFromSession(session))
}
