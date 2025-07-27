package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ThirdPaymentModel = (*customThirdPaymentModel)(nil)

type (
	// ThirdPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPaymentModel.
	ThirdPaymentModel interface {
		thirdPaymentModel
		withSession(session sqlx.Session) ThirdPaymentModel
	}

	customThirdPaymentModel struct {
		*defaultThirdPaymentModel
	}
)

// NewThirdPaymentModel returns a model for the database table.
func NewThirdPaymentModel(conn sqlx.SqlConn) ThirdPaymentModel {
	return &customThirdPaymentModel{
		defaultThirdPaymentModel: newThirdPaymentModel(conn),
	}
}

func (m *customThirdPaymentModel) withSession(session sqlx.Session) ThirdPaymentModel {
	return NewThirdPaymentModel(sqlx.NewSqlConnFromSession(session))
}
