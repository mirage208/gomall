package userInfo

import (
	"net/http"

	"github.com/mirage208/gomall/app/user/api/internal/logic/userInfo"
	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户基本信息展示
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userInfo.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
