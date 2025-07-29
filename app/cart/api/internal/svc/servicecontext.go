package svc

import (
	"github.com/mirage208/gomall/app/cart/api/internal/config"
	"github.com/mirage208/gomall/app/cart/api/internal/middleware"
	"github.com/mirage208/gomall/app/cart/rpc/cartclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	CheckUserState rest.Middleware
	CartRpc        cartclient.Cart
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CheckUserState: middleware.NewCheckUserStateMiddleware().Handle,
		CartRpc:        cartclient.NewCart(zrpc.MustNewClient(c.CartRpc)),
	}
}
