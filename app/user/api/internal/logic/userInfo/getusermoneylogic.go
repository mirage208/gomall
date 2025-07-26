package userInfo

import (
	"context"

	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取账户金额
func NewGetUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMoneyLogic {
	return &GetUserMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMoneyLogic) GetUserMoney(req *types.GetUserMoneyReq) (resp *types.GetUserMoneyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
