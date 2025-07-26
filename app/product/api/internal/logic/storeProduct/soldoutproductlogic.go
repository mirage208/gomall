package storeProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SoldoutProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 下架商品
func NewSoldoutProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoldoutProductLogic {
	return &SoldoutProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SoldoutProductLogic) SoldoutProduct(req *types.SoldoutProductReq) error {
	// todo: add your logic here and delete this line

	return nil
}
