package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryProductListLogic {
	return &CategoryProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryProductListLogic) CategoryProductList(in *product.CategoryProductListReq) (*product.CategoryProductListResp, error) {
	// todo: add your logic here and delete this line

	return &product.CategoryProductListResp{}, nil
}
