package storeProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShelfProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上架商品
func NewShelfProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShelfProductLogic {
	return &ShelfProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShelfProductLogic) ShelfProduct(req *types.ShelfProductReq) error {
	// todo: add your logic here and delete this line

	return nil
}
