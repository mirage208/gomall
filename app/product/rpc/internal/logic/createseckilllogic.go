package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillLogic {
	return &CreateSeckillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSeckillLogic) CreateSeckill(in *product.CreateSeckillReq) (*product.CreateSeckillResp, error) {
	// todo: add your logic here and delete this line

	return &product.CreateSeckillResp{}, nil
}
