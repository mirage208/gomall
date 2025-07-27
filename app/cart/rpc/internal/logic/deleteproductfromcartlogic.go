package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"
	"github.com/mirage208/gomall/app/user/model"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"
	"google.golang.org/grpc/status"

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
	_, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("failed to get user info: %v", err)
		return nil, status.Error(500, "User not found")
	}

	_, err = l.svcCtx.CartModel.FindOne(l.ctx, in.CartId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "Cart item not found")
		}
		return nil, status.Error(500, err.Error())
	}

	if err := l.svcCtx.CartModel.Delete(l.ctx, in.CartId); err != nil {
		l.Logger.Errorf("failed to delete cart item: %v", err)
		return nil, status.Error(500, err.Error())
	}

	return &cart.DeleteProductFromCartResp{}, nil
}
