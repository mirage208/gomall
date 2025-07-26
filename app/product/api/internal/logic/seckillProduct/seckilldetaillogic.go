package seckillProduct

import (
	"context"

	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 秒杀商品详情
func NewSeckillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillDetailLogic {
	return &SeckillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillDetailLogic) SeckillDetail(req *types.GetSeckillDetailReq) (resp *types.GetSeckillDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
