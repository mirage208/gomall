package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSeckillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillLogic {
	return &DeleteSeckillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSeckillLogic) DeleteSeckill(in *product.DeleteSeckillReq) (*product.DeleteSeckillResp, error) {
	// todo: add your logic here and delete this line

	return &product.DeleteSeckillResp{}, nil
}
