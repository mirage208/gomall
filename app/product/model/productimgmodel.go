package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductImgModel = (*customProductImgModel)(nil)

type (
	// ProductImgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductImgModel.
	ProductImgModel interface {
		productImgModel
		withSession(session sqlx.Session) ProductImgModel
	}

	customProductImgModel struct {
		*defaultProductImgModel
	}
)

// NewProductImgModel returns a model for the database table.
func NewProductImgModel(conn sqlx.SqlConn) ProductImgModel {
	return &customProductImgModel{
		defaultProductImgModel: newProductImgModel(conn),
	}
}

func (m *customProductImgModel) withSession(session sqlx.Session) ProductImgModel {
	return NewProductImgModel(sqlx.NewSqlConnFromSession(session))
}
