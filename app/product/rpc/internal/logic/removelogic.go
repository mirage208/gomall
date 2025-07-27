package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/model"
	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	// Check if the product ID is valid
	_, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "product not found")
		}
		l.Logger.Error("Failed to remove product:", err)
		return nil, status.Error(500, "failed to remove product")
	}

	err = l.svcCtx.ProductModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error("Failed to delete product:", err)
		return nil, status.Error(500, "failed to delete product")
	}

	return &product.RemoveResponse{}, nil
}
