package cart

import (
	"context"

	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 购物车商品列表
func NewCartProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductListLogic {
	return &CartProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartProductListLogic) CartProductList(req *types.CartProductListReq) (resp *types.CartProductListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
