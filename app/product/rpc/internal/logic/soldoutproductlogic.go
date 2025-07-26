package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type SoldoutProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSoldoutProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoldoutProductLogic {
	return &SoldoutProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SoldoutProductLogic) SoldoutProduct(in *product.SoldoutProductReq) (*product.SoldoutProductResp, error) {
	// todo: add your logic here and delete this line

	return &product.SoldoutProductResp{}, nil
}
