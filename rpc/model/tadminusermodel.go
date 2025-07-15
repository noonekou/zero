package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TAdminUserModel = (*customTAdminUserModel)(nil)

type (
	// TAdminUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAdminUserModel.
	TAdminUserModel interface {
		tAdminUserModel
		withSession(session sqlx.Session) TAdminUserModel
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

func (m *customTAdminUserModel) withSession(session sqlx.Session) TAdminUserModel {
	return NewTAdminUserModel(sqlx.NewSqlConnFromSession(session))
}
