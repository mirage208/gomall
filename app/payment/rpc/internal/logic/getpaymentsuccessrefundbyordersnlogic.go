package logic

import (
	"context"

	"github.com/mirage208/gomall/app/payment/rpc/internal/svc"
	"github.com/mirage208/gomall/app/payment/rpc/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentSuccessRefundByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentSuccessRefundByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentSuccessRefundByOrderSnLogic {
	return &GetPaymentSuccessRefundByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据订单sn查询流水记录
func (l *GetPaymentSuccessRefundByOrderSnLogic) GetPaymentSuccessRefundByOrderSn(in *payment.GetPaymentSuccessRefundByOrderSnReq) (*payment.GetPaymentSuccessRefundByOrderSnResp, error) {
	// todo: add your logic here and delete this line

	return &payment.GetPaymentSuccessRefundByOrderSnResp{}, nil
}
