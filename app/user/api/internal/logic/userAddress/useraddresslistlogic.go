package userAddress

import (
	"context"

	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户收货地址列表
func NewUserAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressListLogic {
	return &UserAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddressListLogic) UserAddressList() (resp *types.UserAddressListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
