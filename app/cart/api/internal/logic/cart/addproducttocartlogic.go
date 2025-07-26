package cart

import (
	"context"

	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加商品到购物车
func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductToCartLogic) AddProductToCart(req *types.AddProductToCartReq) error {
	// todo: add your logic here and delete this line

	return nil
}
