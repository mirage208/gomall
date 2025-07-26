package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CarouselModel = (*customCarouselModel)(nil)

type (
	// CarouselModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCarouselModel.
	CarouselModel interface {
		carouselModel
		withSession(session sqlx.Session) CarouselModel
	}

	customCarouselModel struct {
		*defaultCarouselModel
	}
)

// NewCarouselModel returns a model for the database table.
func NewCarouselModel(conn sqlx.SqlConn) CarouselModel {
	return &customCarouselModel{
		defaultCarouselModel: newCarouselModel(conn),
	}
}

func (m *customCarouselModel) withSession(session sqlx.Session) CarouselModel {
	return NewCarouselModel(sqlx.NewSqlConnFromSession(session))
}
