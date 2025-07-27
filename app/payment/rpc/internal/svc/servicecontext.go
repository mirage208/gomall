package svc

import (
	"github.com/mirage208/gomall/app/order/rpc/orderclient"
	"github.com/mirage208/gomall/app/payment/model"
	"github.com/mirage208/gomall/app/payment/rpc/internal/config"
	"github.com/mirage208/gomall/app/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PaymentModel model.PayModel

	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		PaymentModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc:     orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
