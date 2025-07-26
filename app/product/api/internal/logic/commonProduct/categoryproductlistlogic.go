package commonProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品分类列表
func NewCategoryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryProductListLogic {
	return &CategoryProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryProductListLogic) CategoryProductList(req *types.CategoryProductListReq) (resp *types.CategoryProductListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
