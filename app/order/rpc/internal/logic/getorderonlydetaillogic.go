package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderOnlyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderOnlyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderOnlyDetailLogic {
	return &GetOrderOnlyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// other
func (l *GetOrderOnlyDetailLogic) GetOrderOnlyDetail(in *order.GetOrderOnlyDetailReq) (*order.GetOrderOnlyDetailResp, error) {
	// todo: add your logic here and delete this line

	return &order.GetOrderOnlyDetailResp{}, nil
}
