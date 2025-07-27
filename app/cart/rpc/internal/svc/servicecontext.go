package svc

import (
	"github.com/mirage208/gomall/app/cart/model"
	"github.com/mirage208/gomall/app/cart/rpc/internal/config"
	"github.com/mirage208/gomall/app/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CartModel model.CartModel

	UserRpcClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		CartModel:     model.NewCartModel(conn, c.CacheRedis),
		UserRpcClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
