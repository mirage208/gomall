package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectProductLogic {
	return &DeleteCollectProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCollectProductLogic) DeleteCollectProduct(in *product.DeleteCollectProductReq) (*product.DeleteCollectProductResp, error) {
	// todo: add your logic here and delete this line

	return &product.DeleteCollectProductResp{}, nil
}
