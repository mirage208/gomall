package logic

import (
	"context"

	"github.com/mirage208/gomall/app/payment/rpc/internal/svc"
	"github.com/mirage208/gomall/app/payment/rpc/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentDetailLogic {
	return &GetPaymentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentDetailLogic) GetPaymentDetail(in *payment.GetPaymentDetailReq) (*payment.GetPaymentDetailResp, error) {
	// todo: add your logic here and delete this line

	return &payment.GetPaymentDetailResp{}, nil
}
