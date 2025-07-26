package userAddress

import (
	"context"

	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新收货地址
func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(req *types.UpdateUserAddressReq) error {
	// todo: add your logic here and delete this line

	return nil
}
