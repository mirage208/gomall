package thirdPayment

import (
	"context"

	"github.com/mirage208/gomall/app/payment/api/internal/svc"
	"github.com/mirage208/gomall/app/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdPaymentwxPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// third payment：wechat pay
func NewThirdPaymentwxPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentwxPayLogic {
	return &ThirdPaymentwxPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentwxPayLogic) ThirdPaymentwxPay(req *types.ThirdPaymentWxPayReq) (resp *types.ThirdPaymentWxPayResp, err error) {
	// todo: add your logic here and delete this line

	return
}
