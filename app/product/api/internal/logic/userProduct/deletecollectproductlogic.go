package userProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除收藏商品
func NewDeleteCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectProductLogic {
	return &DeleteCollectProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectProductLogic) DeleteCollectProduct(req *types.DeleteCollectProductReq) error {
	// todo: add your logic here and delete this line

	return nil
}
