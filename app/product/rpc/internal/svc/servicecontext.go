package svc

import (
	"github.com/mirage208/gomall/app/product/model"
	"github.com/mirage208/gomall/app/product/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ProductModel model.ProductModel
	RDS          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
		RDS:          redis.MustNewRedis(c.RdsConf),
	}
}
