package cart

import (
	"context"

	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductFromCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除购物车商品
func NewDeleteProductFromCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFromCartLogic {
	return &DeleteProductFromCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductFromCartLogic) DeleteProductFromCart(req *types.DeteleProductFromCartReq) error {
	// todo: add your logic here and delete this line

	return nil
}
