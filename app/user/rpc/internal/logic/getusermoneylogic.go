package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMoneyLogic {
	return &GetUserMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserMoneyLogic) GetUserMoney(in *user.GetUserMoneyReq) (*user.GetUserMoneyResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserMoneyResp{}, nil
}
