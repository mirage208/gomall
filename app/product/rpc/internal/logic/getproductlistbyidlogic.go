package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListByIDLogic {
	return &GetProductListByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListByIDLogic) GetProductListByID(in *product.GetProductListByIDReq) (*product.GetProductListByIDResp, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductListByIDResp{}, nil
}
