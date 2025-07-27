package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrderListLogic {
	return &UserOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户订单
func (l *UserOrderListLogic) UserOrderList(in *order.UserOrderListReq) (*order.UserOrderListResp, error) {
	// todo: add your logic here and delete this line

	return &order.UserOrderListResp{}, nil
}
