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

type UpdateProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductDetailLogic {
	return &UpdateProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductDetailLogic) UpdateProductDetail(in *cart.UpdateProductDetailReq) (*cart.UpdateProductDetailResp, error) {
	_, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("failed to get user info: %v", err)
		return nil, status.Error(500, "User not found")
	}

	res, err := l.svcCtx.CartModel.FindOne(l.ctx, in.CartId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "Cart item not found")
		}
		return nil, err
	}

	res.UserId = in.UserId
	res.Count = in.Count
	res.Checked = in.Check
	if err := l.svcCtx.CartModel.Update(l.ctx, res); err != nil {
		l.Logger.Errorf("failed to update cart item: %v", err)
		return nil, status.Error(500, err.Error())
	}

	return &cart.UpdateProductDetailResp{}, nil
}
