package commonProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 首页商品轮播图
func NewCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarouselLogic {
	return &CarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CarouselLogic) Carousel() (resp *types.HomePageCarouselResp, err error) {
	// todo: add your logic here and delete this line

	return
}
