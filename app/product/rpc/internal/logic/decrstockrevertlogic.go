package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockRevertLogic {
	return &DecrStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockRevertLogic) DecrStockRevert(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	// 获取 DTM 子事务屏障，参考：https://dtm.pub/practice/barrier.html
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 执行 DTM 子事务
	if err = barrier.CallWithDB(l.svcCtx.ProductDB, func(tx *sql.Tx) error {
		_, err := l.svcCtx.ProductModel.TxAdjustStock(l.ctx, tx, in.Id, in.Num)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
