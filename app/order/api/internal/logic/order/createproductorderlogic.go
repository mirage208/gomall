package order

import (
	"context"

	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建普通商品订单
func NewCreateProductOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductOrderLogic {
	return &CreateProductOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductOrderLogic) CreateProductOrder(req *types.CreateProductOrderReq) (resp *types.CreateOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
