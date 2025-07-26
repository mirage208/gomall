package payment

import (
	"context"

	"github.com/mirage208/gomall/app/payment/api/internal/svc"
	"github.com/mirage208/gomall/app/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单付款
func NewOrderPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayLogic {
	return &OrderPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderPayLogic) OrderPay(req *types.OrderPayReq) (resp *types.OrderPayResp, err error) {
	// todo: add your logic here and delete this line

	return
}
