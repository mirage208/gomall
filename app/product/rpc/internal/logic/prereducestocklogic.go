package logic

import (
	"context"
	"fmt"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"google.golang.org/grpc/status"

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
	script := `
	local stock = redis.call('get', KEYS[1])
	if not stock then
		return -1
	end
	if tonumber(stock) < tonumber(ARGV[1]) then
		return 0
	end
	redis.call('decrby', KEYS[1], ARGV[1])
	return 1
	`
	productKey := fmt.Sprintf("product:stock:%d", in.ProductId)
	result, err := l.svcCtx.RDS.EvalCtx(l.ctx, script, []string{productKey}, in.Quantity)
	if err != nil {
		l.Logger.Error("Failed to pre-reduce stock:", err)
		return nil, status.Error(500, "failed to pre-reduce stock")
	}
	if result.(int64) == 0 {
		l.Logger.Error("Insufficient stock for product", logx.Field("productId", in.ProductId), logx.Field("quantity", in.Quantity))
		return nil, status.Error(400, "insufficient stock for product")
	}
	if result.(int64) == -1 {
		productInfo, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
		if err != nil {
			l.Logger.Error("Product not found", logx.Field("productId", in.ProductId), logx.Field("error", err))
			return nil, status.Error(404, "product not found")
		}
		if err = l.svcCtx.RDS.SetCtx(l.ctx, productKey, fmt.Sprintf("%d", productInfo.Stock)); err != nil {
			l.Logger.Error("Failed to set stock in Redis", logx.Field("productId", in.ProductId), logx.Field("error", err))
			return nil, status.Error(500, "failed to set stock in Redis")
		}
		// Retry the pre-reduce stock operation
		result, err = l.svcCtx.RDS.EvalCtx(l.ctx, script, []string{productKey}, in.Quantity)
		if err != nil {
			l.Logger.Error("Failed to pre-reduce stock on retry:", err)
			return nil, status.Error(500, "failed to pre-reduce stock on retry")
		}
		if result.(int64) == 0 {
			l.Logger.Error("Insufficient stock for product on retry", logx.Field("productId", in.ProductId), logx.Field("quantity", in.Quantity))
			return nil, status.Error(400, "insufficient stock for product on retry")
		}
	}

	return &product.PreReduceStockResponse{
		ReservationId: fmt.Sprintf("reservation:%d:%d", in.ProductId, in.Quantity),
	}, nil
}
