package order

import (
	"context"

	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建秒杀商品订单
func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(req *types.CreateSeckillOrderReq) (resp *types.CreateOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
