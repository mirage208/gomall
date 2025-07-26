package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillListLogic {
	return &SeckillListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// seckillProduct
func (l *SeckillListLogic) SeckillList(in *product.SeckillListReq) (*product.SeckillListResp, error) {
	// todo: add your logic here and delete this line

	return &product.SeckillListResp{}, nil
}
