package svc

import (
	"database/sql"

	"github.com/mirage208/gomall/app/product/model"
	"github.com/mirage208/gomall/app/product/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ProductDB    *sql.DB
	ProductModel model.ProductModel
	RDS          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	rawDB, err := conn.RawDB()
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:       c,
		ProductDB:    rawDB,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
		RDS:          redis.MustNewRedis(c.RdsConf),
	}
}
