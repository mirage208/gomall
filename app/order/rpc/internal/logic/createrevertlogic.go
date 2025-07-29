package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/mirage208/gomall/app/order/rpc/internal/svc"
	"github.com/mirage208/gomall/app/order/rpc/pb/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRevertLogic) CreateRevert(in *order.CreateRequest) (*order.CreateResponse, error) {
	lastOrder, err := l.svcCtx.OrderModel.FindLastOneByUidPid(l.ctx, in.Uid, in.Pid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Logger.Error("Failed to get barrier from gRPC context", logx.Field("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := barrier.CallWithDB(l.svcCtx.OrderDB, func(tx *sql.Tx) error {
		lastOrder.Status = -1 // Revert the order status
		if err := l.svcCtx.OrderModel.Update(l.ctx, lastOrder); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &order.CreateResponse{}, nil
}
