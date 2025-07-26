package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(in *order.CreateSeckillOrderReq) (*order.CreateOrderResp, error) {
	// todo: add your logic here and delete this line

	return &order.CreateOrderResp{}, nil
}
