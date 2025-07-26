package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShelfProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShelfProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShelfProductLogic {
	return &ShelfProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShelfProductLogic) ShelfProduct(in *product.ShelfProductReq) (*product.ShelfProductResp, error) {
	// todo: add your logic here and delete this line

	return &product.ShelfProductResp{}, nil
}
