package logic

import (
	"context"
	"database/sql"

	// 下面这行导入gozero的dtm驱动
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockLogic) DecrStock(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {

	// 获取 DTM 子事务屏障，参考：https://dtm.pub/practice/barrier.html
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 执行 DTM 子事务
	if err = barrier.CallWithDB(l.svcCtx.ProductDB, func(tx *sql.Tx) error {
		// decrease the stock of the product
		res, err := l.svcCtx.ProductModel.AdjustStock(l.ctx, in.Id, -in.Num)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		affected, err := res.RowsAffected()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		// 受影响行数为 0 时，库存不足导致扣件失败
		// 此时不需要重试，直接告诉 DTM 事务失败
		if affected <= 0 {
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return &product.DecrStockResponse{}, nil
}
