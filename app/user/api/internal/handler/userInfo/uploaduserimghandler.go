package userInfo

import (
	"net/http"

	"github.com/mirage208/gomall/app/user/api/internal/logic/userInfo"
	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上传头像
func UploadUserImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadUserImgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userInfo.NewUploadUserImgLogic(r.Context(), svcCtx)
		err := l.UploadUserImg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
