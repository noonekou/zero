package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TUserRoleModel = (*customTUserRoleModel)(nil)

type (
	// TUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserRoleModel.
	TUserRoleModel interface {
		tUserRoleModel
		withSession(session sqlx.Session) TUserRoleModel
	}

	customTUserRoleModel struct {
		*defaultTUserRoleModel
	}
)

// NewTUserRoleModel returns a model for the database table.
func NewTUserRoleModel(conn sqlx.SqlConn) TUserRoleModel {
	return &customTUserRoleModel{
		defaultTUserRoleModel: newTUserRoleModel(conn),
	}
}

func (m *customTUserRoleModel) withSession(session sqlx.Session) TUserRoleModel {
	return NewTUserRoleModel(sqlx.NewSqlConnFromSession(session))
}
