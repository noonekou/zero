package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TApisModel = (*customTApisModel)(nil)

type (
	// TApisModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTApisModel.
	TApisModel interface {
		tApisModel
		withSession(session sqlx.Session) TApisModel
	}

	customTApisModel struct {
		*defaultTApisModel
	}
)

// NewTApisModel returns a model for the database table.
func NewTApisModel(conn sqlx.SqlConn) TApisModel {
	return &customTApisModel{
		defaultTApisModel: newTApisModel(conn),
	}
}

func (m *customTApisModel) withSession(session sqlx.Session) TApisModel {
	return NewTApisModel(sqlx.NewSqlConnFromSession(session))
}
