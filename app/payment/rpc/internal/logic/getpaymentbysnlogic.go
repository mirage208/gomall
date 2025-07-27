package logic

import (
	"context"

	"github.com/mirage208/gomall/app/payment/rpc/internal/svc"
	"github.com/mirage208/gomall/app/payment/rpc/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentBySnLogic {
	return &GetPaymentBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据sn查询流水记录
func (l *GetPaymentBySnLogic) GetPaymentBySn(in *payment.GetPaymentBySnReq) (*payment.GetPaymentBySnResp, error) {
	// todo: add your logic here and delete this line

	return &payment.GetPaymentBySnResp{}, nil
}
