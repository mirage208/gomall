package logic

import (
	"context"

	"github.com/mirage208/gomall/app/payment/rpc/internal/svc"
	"github.com/mirage208/gomall/app/payment/rpc/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeStateLogic {
	return &UpdateTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新交易状态
func (l *UpdateTradeStateLogic) UpdateTradeState(in *payment.UpdateTradeStateReq) (*payment.UpdateTradeStateResp, error) {
	// todo: add your logic here and delete this line

	return &payment.UpdateTradeStateResp{}, nil
}
