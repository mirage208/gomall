package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillDetailLogic {
	return &SeckillDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SeckillDetailLogic) SeckillDetail(in *product.SeckillDetailReq) (*product.SeckillDetailResp, error) {
	// todo: add your logic here and delete this line

	return &product.SeckillDetailResp{}, nil
}
