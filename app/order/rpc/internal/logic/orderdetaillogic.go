package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单详情
func (l *OrderDetailLogic) OrderDetail(in *order.OrderDetailReq) (*order.OrderDetailResp, error) {
	// todo: add your logic here and delete this line

	return &order.OrderDetailResp{}, nil
}
