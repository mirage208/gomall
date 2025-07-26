package commonProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品详情信息
func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ProductDetailReq) (resp *types.ProductDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
