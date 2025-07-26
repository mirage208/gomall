package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductCommentModel = (*customProductCommentModel)(nil)

type (
	// ProductCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCommentModel.
	ProductCommentModel interface {
		productCommentModel
		withSession(session sqlx.Session) ProductCommentModel
	}

	customProductCommentModel struct {
		*defaultProductCommentModel
	}
)

// NewProductCommentModel returns a model for the database table.
func NewProductCommentModel(conn sqlx.SqlConn) ProductCommentModel {
	return &customProductCommentModel{
		defaultProductCommentModel: newProductCommentModel(conn),
	}
}

func (m *customProductCommentModel) withSession(session sqlx.Session) ProductCommentModel {
	return NewProductCommentModel(sqlx.NewSqlConnFromSession(session))
}
