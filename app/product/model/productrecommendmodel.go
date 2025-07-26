package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductRecommendModel = (*customProductRecommendModel)(nil)

type (
	// ProductRecommendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductRecommendModel.
	ProductRecommendModel interface {
		productRecommendModel
		withSession(session sqlx.Session) ProductRecommendModel
	}

	customProductRecommendModel struct {
		*defaultProductRecommendModel
	}
)

// NewProductRecommendModel returns a model for the database table.
func NewProductRecommendModel(conn sqlx.SqlConn) ProductRecommendModel {
	return &customProductRecommendModel{
		defaultProductRecommendModel: newProductRecommendModel(conn),
	}
}

func (m *customProductRecommendModel) withSession(session sqlx.Session) ProductRecommendModel {
	return NewProductRecommendModel(sqlx.NewSqlConnFromSession(session))
}
