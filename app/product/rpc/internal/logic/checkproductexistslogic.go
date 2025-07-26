package logic

import (
	"context"

	"github.com/mirage208/gomall/app/product/rpc/internal/svc"
	"github.com/mirage208/gomall/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckProductExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckProductExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductExistsLogic {
	return &CheckProductExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// others
func (l *CheckProductExistsLogic) CheckProductExists(in *product.CheckProductExistsReq) (*product.CheckProductExistsResp, error) {
	// todo: add your logic here and delete this line

	return &product.CheckProductExistsResp{}, nil
}
