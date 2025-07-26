package userAddress

import (
	"context"

	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加收货地址
func NewCreateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserAddressLogic {
	return &CreateUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserAddressLogic) CreateUserAddress(req *types.CreateUserAddressReq) error {
	// todo: add your logic here and delete this line

	return nil
}
