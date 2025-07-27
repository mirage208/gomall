package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/model"
	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

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
	// check if the user exists
	_, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("failed to get user info: %v", err)
		return nil, err
	}

	// check if the product exists
	_, err = l.svcCtx.ProductRpcClient.Detail(l.ctx, &product.DetailRequest{
		Id: in.ProductId,
	})
	if err != nil {
		l.Logger.Errorf("failed to get product info: %v", err)
		return nil, err
	}

	// add product to cart logic
	newCartItem := &model.Cart{
		UserId:    in.UserId,
		ProductId: in.ProductId,
		Count:     in.Count,
		Checked:   1, // default to checked
	}
	_, err = l.svcCtx.CartModel.Insert(l.ctx, newCartItem)
	if err != nil {
		l.Logger.Errorf("failed to insert cart item: %v", err)
		return nil, err
	}

	return &cart.AddProductToCartResp{}, nil
}
