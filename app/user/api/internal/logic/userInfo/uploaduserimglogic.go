package userInfo

import (
	"context"

	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传头像
func NewUploadUserImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserImgLogic {
	return &UploadUserImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadUserImgLogic) UploadUserImg(req *types.UploadUserImgReq) error {
	// todo: add your logic here and delete this line

	return nil
}
