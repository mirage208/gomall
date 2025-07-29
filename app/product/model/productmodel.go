package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		AdjustStock(ctx context.Context, productId int64, quantity int64) (sql.Result, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

// AdjustStock updates the stock of a product by the given quantity.
// If quantity is negative, it will decrease the stock; if positive, it will increase the stock.
func (m *customProductModel) AdjustStock(ctx context.Context, productId int64, quantity int64) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, productId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set stock = stock + ? where `id` = ? and stock >= -?", m.table)
		return conn.ExecCtx(ctx, query, quantity, productId, quantity)
	}, productIdKey)
	return ret, err
}
