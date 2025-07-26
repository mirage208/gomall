package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserAndAddressExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserAndAddressExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserAndAddressExistsLogic {
	return &CheckUserAndAddressExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserAndAddressExistsLogic) CheckUserAndAddressExists(in *user.CheckUserAndAddressExistsReq) (*user.CheckUserAndAddressExistsResp, error) {
	// todo: add your logic here and delete this line

	return &user.CheckUserAndAddressExistsResp{}, nil
}
