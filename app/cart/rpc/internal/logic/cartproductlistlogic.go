package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductListLogic {
	return &CartProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartProductListLogic) CartProductList(in *cart.CartProductListReq) (*cart.CartProductListResp, error) {
	// todo: add your logic here and delete this line

	return &cart.CartProductListResp{}, nil
}
