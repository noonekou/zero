package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TResourceModel = (*customTResourceModel)(nil)

type (
	// TResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTResourceModel.
	TResourceModel interface {
		tResourceModel
		withSession(session sqlx.Session) TResourceModel
	}

	customTResourceModel struct {
		*defaultTResourceModel
	}
)

// NewTResourceModel returns a model for the database table.
func NewTResourceModel(conn sqlx.SqlConn) TResourceModel {
	return &customTResourceModel{
		defaultTResourceModel: newTResourceModel(conn),
	}
}

func (m *customTResourceModel) withSession(session sqlx.Session) TResourceModel {
	return NewTResourceModel(sqlx.NewSqlConnFromSession(session))
}
