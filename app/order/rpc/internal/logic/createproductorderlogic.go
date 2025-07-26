package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductOrderLogic {
	return &CreateProductOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductOrderLogic) CreateProductOrder(in *order.CreateProductOrderReq) (*order.CreateOrderResp, error) {
	// todo: add your logic here and delete this line

	return &order.CreateOrderResp{}, nil
}
