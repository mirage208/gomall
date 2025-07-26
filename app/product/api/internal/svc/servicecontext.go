package svc

import (
	"github.com/mirage208/gomall/app/product/api/internal/config"
	"github.com/mirage208/gomall/app/product/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	CheckStoreState rest.Middleware
	CheckUserState  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CheckStoreState: middleware.NewCheckStoreStateMiddleware().Handle,
		CheckUserState:  middleware.NewCheckUserStateMiddleware().Handle,
	}
}
