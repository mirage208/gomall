package svc

import (
	"github.com/mirage208/gomall/app/order/api/internal/config"
	"github.com/mirage208/gomall/app/order/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	CheckUserState rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CheckUserState: middleware.NewCheckUserStateMiddleware().Handle,
	}
}
