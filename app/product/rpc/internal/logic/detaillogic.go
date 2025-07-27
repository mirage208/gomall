package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/model"
	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	// Check if the product ID is valid
	productInfo, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "product not found")
		}
		l.Logger.Error("Failed to fetch product details:", err)
		return nil, status.Error(500, "failed to fetch product details")
	}

	return &product.DetailResponse{
		Id:     productInfo.Id,
		Name:   productInfo.Name,
		Desc:   productInfo.Desc,
		Stock:  productInfo.Stock,
		Amount: productInfo.Amount,
		Status: productInfo.Status,
	}, nil
}
