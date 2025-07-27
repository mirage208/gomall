package logic

import (
	"context"

	"github.com/mirage208/gomall/app/user/model"
	"github.com/mirage208/gomall/app/user/rpc/internal/svc"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(in *user.UserInfoUpdateRequest) (*user.UserInfoUpdateResponse, error) {
	// find user by ID
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// update user info
	if in.Name != nil {
		userInfo.Name = *in.Name
	}
	if in.Gender != nil {
		userInfo.Gender = *in.Gender
	}
	if in.Mobile != nil {
		userInfo.Mobile = *in.Mobile
	}
	if err = l.svcCtx.UserModel.Update(l.ctx, userInfo); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoUpdateResponse{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Gender: userInfo.Gender,
		Mobile: userInfo.Mobile,
	}, nil
}
