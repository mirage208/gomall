package cart

import (
	"context"

	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改购物车商品的信息
func NewUpdateProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductDetailLogic {
	return &UpdateProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductDetailLogic) UpdateProductDetail(req *types.UpdateProductDetailReq) (resp *types.UpdateProductDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
