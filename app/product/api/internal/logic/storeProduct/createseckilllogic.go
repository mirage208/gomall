package storeProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加秒杀商品
func NewCreateSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillLogic {
	return &CreateSeckillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillLogic) CreateSeckill(req *types.CreateSeckillProductReq) error {
	// todo: add your logic here and delete this line

	return nil
}
