package userProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCollectProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加收藏商品
func NewCreateCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCollectProductLogic {
	return &CreateCollectProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCollectProductLogic) CreateCollectProduct(req *types.CreateCollectProductReq) error {
	// todo: add your logic here and delete this line

	return nil
}
