package handler

import (
	"net/http"

	"github.com/mirage208/gomall/app/user/api/internal/logic"
	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登陆
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
