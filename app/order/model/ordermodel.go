package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		TxInsert(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error)
		TxUpdate(ctx context.Context, tx *sql.Tx, data *Order) error
		FindAllByUid(ctx context.Context, uid int64) ([]*Order, error)
		FindLastOneByUidPid(ctx context.Context, uid, pid int64) (*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c, opts...),
	}
}

func (m *customOrderModel) FindAllByUid(ctx context.Context, uid int64) ([]*Order, error) {
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid` = ? order by `create_time` desc", orderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOrderModel) TxInsert(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := tx.ExecContext(ctx, query, data.Uid, data.Pid, data.CreateTime)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *customOrderModel) TxUpdate(ctx context.Context, tx *sql.Tx, data *Order) error {
	query := fmt.Sprintf("update %s set `uid` = ?, `pid` = ?, `amount` = ?, `status` = ? where `id` = ?", m.table)
	_, err := tx.ExecContext(ctx, query, data.Uid, data.Pid, data.Amount, data.Status, data.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *customOrderModel) FindLastOneByUidPid(ctx context.Context, uid, pid int64) (*Order, error) {
	var resp Order

	query := fmt.Sprintf("select %s from %s where `uid` = ? and `pid` = ? order by `create_time` desc limit 1", orderRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, uid, pid)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
