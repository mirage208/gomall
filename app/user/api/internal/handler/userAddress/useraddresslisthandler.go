package userAddress

import (
	"net/http"

	"github.com/mirage208/gomall/app/user/api/internal/logic/userAddress"
	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户收货地址列表
func UserAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userAddress.NewUserAddressListLogic(r.Context(), svcCtx)
		resp, err := l.UserAddressList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
