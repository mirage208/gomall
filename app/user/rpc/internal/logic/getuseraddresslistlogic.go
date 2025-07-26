package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressListLogic {
	return &GetUserAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// userAddress
func (l *GetUserAddressListLogic) GetUserAddressList(in *user.GetUserAddressListReq) (*user.GetUserAddressListResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserAddressListResp{}, nil
}
