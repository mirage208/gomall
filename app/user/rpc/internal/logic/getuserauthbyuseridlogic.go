package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *user.GetUserAuthByUserIdReq) (*user.GetUserAuthyUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserAuthyUserIdResp{}, nil
}
