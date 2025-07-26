package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentListLogic {
	return &ProductCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCommentListLogic) ProductCommentList(in *product.ProductCommentListReq) (*product.ProductCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &product.ProductCommentListResp{}, nil
}
