package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMoneyLogic {
	return &UpdateUserMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserMoneyLogic) UpdateUserMoney(in *user.UpdateUserMoneyReq) (*user.UpdateUserMoneyResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateUserMoneyResp{}, nil
}
