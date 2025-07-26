package userAddress

import (
	"context"

	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除收货地址
func NewDeleteUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserAddressLogic {
	return &DeleteUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserAddressLogic) DeleteUserAddress(req *types.DeleteUserAddressReq) error {
	// todo: add your logic here and delete this line

	return nil
}
