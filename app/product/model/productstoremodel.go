package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductStoreModel = (*customProductStoreModel)(nil)

type (
	// ProductStoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductStoreModel.
	ProductStoreModel interface {
		productStoreModel
		withSession(session sqlx.Session) ProductStoreModel
	}

	customProductStoreModel struct {
		*defaultProductStoreModel
	}
)

// NewProductStoreModel returns a model for the database table.
func NewProductStoreModel(conn sqlx.SqlConn) ProductStoreModel {
	return &customProductStoreModel{
		defaultProductStoreModel: newProductStoreModel(conn),
	}
}

func (m *customProductStoreModel) withSession(session sqlx.Session) ProductStoreModel {
	return NewProductStoreModel(sqlx.NewSqlConnFromSession(session))
}
