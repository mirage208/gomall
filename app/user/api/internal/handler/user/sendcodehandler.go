package user

import (
	"net/http"

	"github.com/mirage208/gomall/app/user/api/internal/logic/user"
	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 邮箱验证码发送
func SendCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendCodeLogic(r.Context(), svcCtx)
		err := l.SendCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
