package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(in *user.UpdateUserAddressReq) (*user.UpdateUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateUserAddressResp{}, nil
}
