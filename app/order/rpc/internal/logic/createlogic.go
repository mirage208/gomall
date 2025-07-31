package logic

import (
	"context"
	"database/sql"

	"github.com/mirage208/gomall/app/order/model"
	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dtm-labs/client/dtmgrpc"
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
		return nil, status.Error(codes.Aborted, "User not found")
	}

	// check if the product exists
	productInfo, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{Id: in.Pid})
	if err != nil {
		l.Logger.Error("Product not found", logx.Field("productId", in.Pid), logx.Field("error", err))
		return nil, status.Error(codes.Aborted, "Product not found")
	}

	// check if the product is in stock and pre-reduce stock
	if productInfo.Stock <= 0 {
		l.Logger.Error("Product out of stock", logx.Field("productId", in.Pid))
		return nil, status.Error(codes.Aborted, "Product out of stock")
	}

	// get the DTM barrier from the gRPC context
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Logger.Error("Failed to get barrier from gRPC context", logx.Field("error", err))
		return nil, status.Error(codes.Internal, "failed to get barrier")
	}

	newOrder := &model.Order{
		Uid: in.Uid,
		Pid: in.Pid,
	}
	// use the barrier to ensure atomicity
	if err := barrier.CallWithDB(l.svcCtx.OrderDB, func(tx *sql.Tx) error {
		res, err := l.svcCtx.OrderModel.TxInsert(l.ctx, tx, newOrder)
		if err != nil {
			l.Logger.Errorf("Failed to create order: %v", err)
			return status.Error(codes.Internal, err.Error())
		}
		newOrder.Id, err = res.LastInsertId()
		if err != nil {
			l.Logger.Errorf("Failed to get last insert ID: %v", err)
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &order.CreateResponse{Id: newOrder.Id}, nil
}
