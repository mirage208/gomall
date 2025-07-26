package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserAddressLogic {
	return &DeleteUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserAddressLogic) DeleteUserAddress(in *user.DeleteUserAddressReq) (*user.DeleteUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteUserAddressResp{}, nil
}
