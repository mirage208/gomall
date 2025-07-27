package thirdPayment

import (
	"context"

	"github.com/mirage208/gomall/app/payment/api/internal/svc"
	"github.com/mirage208/gomall/app/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdPaymentWxPayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// third payment：wechat pay callback
func NewThirdPaymentWxPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentWxPayCallbackLogic {
	return &ThirdPaymentWxPayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentWxPayCallbackLogic) ThirdPaymentWxPayCallback(req *types.ThirdPaymentWxPayCallbackReq) (resp *types.ThirdPaymentWxPayCallbackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
