package userProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 收藏商品列表
func NewCollectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectProductListLogic {
	return &CollectProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectProductListLogic) CollectProductList() (resp *types.CollectProductListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
