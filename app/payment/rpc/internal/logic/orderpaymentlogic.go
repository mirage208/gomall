package logic

import (
	"context"

	"github.com/mirage208/gomall/app/payment/rpc/internal/svc"
	"github.com/mirage208/gomall/app/payment/rpc/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentLogic {
	return &OrderPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPaymentLogic) OrderPayment(in *payment.OrderPaymentReq) (*payment.OrderPaymentResp, error) {
	// todo: add your logic here and delete this line

	return &payment.OrderPaymentResp{}, nil
}
