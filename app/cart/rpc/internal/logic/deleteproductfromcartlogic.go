package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductFromCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductFromCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFromCartLogic {
	return &DeleteProductFromCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductFromCartLogic) DeleteProductFromCart(in *cart.DeleteProductFromCartReq) (*cart.DeleteProductFromCartResp, error) {
	// todo: add your logic here and delete this line

	return &cart.DeleteProductFromCartResp{}, nil
}
