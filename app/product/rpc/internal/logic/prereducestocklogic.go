package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreReduceStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreReduceStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreReduceStockLogic {
	return &PreReduceStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreReduceStockLogic) PreReduceStock(in *product.PreReduceStockRequest) (*product.PreReduceStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.PreReduceStockResponse{}, nil
}
