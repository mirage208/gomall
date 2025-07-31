package svc

import (
	"github.com/mirage208/gomall/app/order/api/internal/config"
	"github.com/mirage208/gomall/app/order/rpc/orderclient"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"github.com/mirage208/gomall/app/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   orderclient.Order
	ProductRpc product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
