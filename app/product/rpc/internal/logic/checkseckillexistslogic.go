package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSeckillExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckSeckillExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSeckillExistsLogic {
	return &CheckSeckillExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckSeckillExistsLogic) CheckSeckillExists(in *product.CheckSeckillExistsReq) (*product.CheckSeckillExistsResp, error) {
	// todo: add your logic here and delete this line

	return &product.CheckSeckillExistsResp{}, nil
}
