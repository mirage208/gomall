package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserAddressLogic {
	return &CreateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserAddressLogic) CreateUserAddress(in *user.CreateUserAddressReq) (*user.CreateUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &user.CreateUserAddressResp{}, nil
}
