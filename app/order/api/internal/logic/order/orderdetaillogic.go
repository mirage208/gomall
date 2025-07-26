package order

import (
	"context"

	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单详情
func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.GetOrderDetailReq) (resp *types.GetOrderDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
