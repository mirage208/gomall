package cart

import (
	"net/http"

	"github.com/mirage208/gomall/app/cart/api/internal/logic/cart"
	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除购物车商品
func DeleteProductFromCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductFromCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewDeleteProductFromCartLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductFromCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
