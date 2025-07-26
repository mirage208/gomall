package storeProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSeckillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除秒杀商品
func NewDeleteSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillLogic {
	return &DeleteSeckillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSeckillLogic) DeleteSeckill(req *types.DeleteSeckillProductReq) error {
	// todo: add your logic here and delete this line

	return nil
}
