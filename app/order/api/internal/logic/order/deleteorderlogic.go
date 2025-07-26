package order

import (
	"context"

	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除订单信息
func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrderLogic) DeleteOrder(req *types.DeleteOrderReq) error {
	// todo: add your logic here and delete this line

	return nil
}
