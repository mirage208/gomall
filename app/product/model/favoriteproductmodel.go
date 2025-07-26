package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FavoriteProductModel = (*customFavoriteProductModel)(nil)

type (
	// FavoriteProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoriteProductModel.
	FavoriteProductModel interface {
		favoriteProductModel
		withSession(session sqlx.Session) FavoriteProductModel
	}

	customFavoriteProductModel struct {
		*defaultFavoriteProductModel
	}
)

// NewFavoriteProductModel returns a model for the database table.
func NewFavoriteProductModel(conn sqlx.SqlConn) FavoriteProductModel {
	return &customFavoriteProductModel{
		defaultFavoriteProductModel: newFavoriteProductModel(conn),
	}
}

func (m *customFavoriteProductModel) withSession(session sqlx.Session) FavoriteProductModel {
	return NewFavoriteProductModel(sqlx.NewSqlConnFromSession(session))
}
