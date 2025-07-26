package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SeckillProductModel = (*customSeckillProductModel)(nil)

type (
	// SeckillProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSeckillProductModel.
	SeckillProductModel interface {
		seckillProductModel
		withSession(session sqlx.Session) SeckillProductModel
	}

	customSeckillProductModel struct {
		*defaultSeckillProductModel
	}
)

// NewSeckillProductModel returns a model for the database table.
func NewSeckillProductModel(conn sqlx.SqlConn) SeckillProductModel {
	return &customSeckillProductModel{
		defaultSeckillProductModel: newSeckillProductModel(conn),
	}
}

func (m *customSeckillProductModel) withSession(session sqlx.Session) SeckillProductModel {
	return NewSeckillProductModel(sqlx.NewSqlConnFromSession(session))
}
