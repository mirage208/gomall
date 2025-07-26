package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCollectProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCollectProductLogic {
	return &CreateCollectProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// userProduct
func (l *CreateCollectProductLogic) CreateCollectProduct(in *product.CreateCollectProductReq) (*product.CreateCollectProductResp, error) {
	// todo: add your logic here and delete this line

	return &product.CreateCollectProductResp{}, nil
}
