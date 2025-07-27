package logic

import (
	"context"

	"github.com/mirage208/gomall/app/cart/rpc/internal/svc"
	"github.com/mirage208/gomall/app/cart/rpc/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductDetailLogic {
	return &UpdateProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductDetailLogic) UpdateProductDetail(in *cart.UpdateProductDetailReq) (*cart.UpdateProductDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cart.UpdateProductDetailResp{}, nil
}
