package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserExistsLogic {
	return &CheckUserExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// other
func (l *CheckUserExistsLogic) CheckUserExists(in *user.CheckUserExistsReq) (*user.CheckUserExistsResp, error) {
	// todo: add your logic here and delete this line

	return &user.CheckUserExistsResp{}, nil
}
