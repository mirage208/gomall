package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderTradeStateLogic {
	return &UpdateOrderTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新订单状态
func (l *UpdateOrderTradeStateLogic) UpdateOrderTradeState(in *order.UpdateOrderTradeStateReq) (*order.UpdateOrderTradeStateResp, error) {
	// todo: add your logic here and delete this line

	return &order.UpdateOrderTradeStateResp{}, nil
}
