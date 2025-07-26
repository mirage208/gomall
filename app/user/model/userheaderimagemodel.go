package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserHeaderImageModel = (*customUserHeaderImageModel)(nil)

type (
	// UserHeaderImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserHeaderImageModel.
	UserHeaderImageModel interface {
		userHeaderImageModel
		withSession(session sqlx.Session) UserHeaderImageModel
	}

	customUserHeaderImageModel struct {
		*defaultUserHeaderImageModel
	}
)

// NewUserHeaderImageModel returns a model for the database table.
func NewUserHeaderImageModel(conn sqlx.SqlConn) UserHeaderImageModel {
	return &customUserHeaderImageModel{
		defaultUserHeaderImageModel: newUserHeaderImageModel(conn),
	}
}

func (m *customUserHeaderImageModel) withSession(session sqlx.Session) UserHeaderImageModel {
	return NewUserHeaderImageModel(sqlx.NewSqlConnFromSession(session))
}
