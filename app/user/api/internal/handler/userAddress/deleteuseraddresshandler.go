package userAddress

import (
	"net/http"

	"github.com/mirage208/gomall/app/user/api/internal/logic/userAddress"
	"github.com/mirage208/gomall/app/user/api/internal/svc"
	"github.com/mirage208/gomall/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除收货地址
func DeleteUserAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userAddress.NewDeleteUserAddressLogic(r.Context(), svcCtx)
		err := l.DeleteUserAddress(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
