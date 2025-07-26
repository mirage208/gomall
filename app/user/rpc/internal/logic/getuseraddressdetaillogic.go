package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressDetailLogic {
	return &GetUserAddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAddressDetailLogic) GetUserAddressDetail(in *user.GetUserAddressDetailReq) (*user.GetUserAddressDetailResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserAddressDetailResp{}, nil
}
