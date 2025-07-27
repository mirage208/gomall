package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/model"
	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// check if the product ID is valid
	productInfo, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "product not found")
		}
		l.Logger.Error("Failed to fetch product details:", err)
		return nil, status.Error(500, "failed to fetch product details")
	}

	// Update product details
	if in.Name != nil {
		productInfo.Name = *in.Name
	}
	if in.Desc != nil {
		productInfo.Desc = *in.Desc
	}
	if in.Stock != nil {
		productInfo.Stock = *in.Stock
	}
	if in.Amount != nil {
		productInfo.Amount = *in.Amount
	}
	if in.Status != nil {
		productInfo.Status = *in.Status
	}
	err = l.svcCtx.ProductModel.Update(l.ctx, productInfo)
	if err != nil {
		l.Logger.Error("Failed to update product:", err)
		return nil, status.Error(500, "failed to update product")
	}

	return &product.UpdateResponse{}, nil
}
