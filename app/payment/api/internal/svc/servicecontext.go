package svc

import (
	"github.com/mirage208/gomall/app/payment/api/internal/config"
	"github.com/mirage208/gomall/app/payment/rpc/paymentclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	PaymentRpc paymentclient.Payment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		PaymentRpc: paymentclient.NewPayment(zrpc.MustNewClient(c.PaymentRpc)),
	}
}
