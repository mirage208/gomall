package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

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
	_, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("failed to get user info: %v", err)
		return nil, err
	}
	// TODO: implement cart product list logic
	list, err := l.svcCtx.CartModel.FindAllByUserId(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to get cart items: %v", err)
		return nil, err
	}
	cartProducts := make([]*cart.CartProduct, 0, len(list))
	for _, item := range list {
		cartProducts = append(cartProducts, &cart.CartProduct{
			Id:        item.Id,
			ProductId: item.ProductId,
			Count:     item.Count,
			Checked:   item.Checked,
		})
	}
	return &cart.CartProductListResp{
		CartProducts: cartProducts,
	}, nil
}
