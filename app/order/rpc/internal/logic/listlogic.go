package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/model"
	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	// check if the user exists
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	list, err := l.svcCtx.OrderModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// Convert the list of orders to the response format
	orderList := make([]*order.DetailResponse, len(list))
	for i, orderItem := range list {
		orderList[i] = &order.DetailResponse{
			Id:     orderItem.Id,
			Uid:    orderItem.Uid,
			Pid:    orderItem.Pid,
			Amount: orderItem.Amount,
			Status: orderItem.Status,
		}
	}

	return &order.ListResponse{
		Data: orderList,
	}, nil
}
