package commonProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品搜索
func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProductLogic) SearchProduct(req *types.SearchProductReq) (resp *types.SearchProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
