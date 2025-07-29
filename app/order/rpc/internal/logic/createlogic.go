package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/model"
	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// check if the user exists
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		l.Logger.Error("User not found", logx.Field("userId", in.Uid), logx.Field("error", err))
		return nil, err
	}

	// check if the product exists
	productInfo, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{Id: in.Pid})
	if err != nil {
		l.Logger.Error("Product not found", logx.Field("productId", in.Pid), logx.Field("error", err))
		return nil, err
	}

	// check if the product is in stock and pre-reduce stock
	if productInfo.Stock <= 0 {
		l.Logger.Error("Product out of stock", logx.Field("productId", in.Pid))
		return nil, status.Error(400, "Product out of stock")
	}
	_, err = l.svcCtx.ProductRpc.PreReduceStock(l.ctx, &product.PreReduceStockRequest{
		ProductId: productInfo.Id,
		Quantity:  1,
	})
	if err != nil {
		l.Logger.Error("Failed to pre-reduce stock", logx.Field("productId", in.Pid), logx.Field("error", err))
		return nil, status.Error(500, "failed to pre-reduce stock")
	}

	// create the order
	newOrder := &model.Order{
		Uid: in.Uid,
		Pid: in.Pid,
	}
	res, err := l.svcCtx.OrderModel.Insert(l.ctx, newOrder)
	if err != nil {
		l.Logger.Errorf("Failed to create order: %v", err)
		return nil, status.Error(500, err.Error())
	}
	newOrder.Id, err = res.LastInsertId()
	if err != nil {
		l.Logger.Errorf("Failed to get last insert ID: %v", err)
		return nil, status.Error(500, err.Error())
	}
	// Update product stock
	newStock := productInfo.Stock - 1
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:    productInfo.Id,
		Stock: &newStock,
	})
	if err != nil {
		l.Logger.Errorf("Failed to update product stock: %v", err)
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{Id: newOrder.Id}, nil
}
