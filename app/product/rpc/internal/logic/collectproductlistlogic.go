package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectProductListLogic {
	return &CollectProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectProductListLogic) CollectProductList(in *product.CollectProductListReq) (*product.CollectProductListResp, error) {
	// todo: add your logic here and delete this line

	return &product.CollectProductListResp{}, nil
}
