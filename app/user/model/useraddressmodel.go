package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserAddressModel = (*customUserAddressModel)(nil)

type (
	// UserAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAddressModel.
	UserAddressModel interface {
		userAddressModel
		withSession(session sqlx.Session) UserAddressModel
	}

	customUserAddressModel struct {
		*defaultUserAddressModel
	}
)

// NewUserAddressModel returns a model for the database table.
func NewUserAddressModel(conn sqlx.SqlConn) UserAddressModel {
	return &customUserAddressModel{
		defaultUserAddressModel: newUserAddressModel(conn),
	}
}

func (m *customUserAddressModel) withSession(session sqlx.Session) UserAddressModel {
	return NewUserAddressModel(sqlx.NewSqlConnFromSession(session))
}
