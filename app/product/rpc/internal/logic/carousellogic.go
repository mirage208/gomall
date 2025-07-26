package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarouselLogic {
	return &CarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// commonProduct
func (l *CarouselLogic) Carousel(in *product.CarouselReq) (*product.CarouselResp, error) {
	// todo: add your logic here and delete this line

	return &product.CarouselResp{}, nil
}
