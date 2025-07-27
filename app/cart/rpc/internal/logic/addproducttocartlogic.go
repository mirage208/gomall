package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductToCartLogic) AddProductToCart(in *cart.AddProductToCartReq) (*cart.AddProductToCartResp, error) {
	// todo: add your logic here and delete this line

	return &cart.AddProductToCartResp{}, nil
}
