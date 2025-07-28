package cart

import (
	"net/http"

	"github.com/mirage208/gomall/app/cart/api/internal/logic/cart"
	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加商品到购物车
func AddProductToCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddProductToCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewAddProductToCartLogic(r.Context(), svcCtx)
		resp, err := l.AddProductToCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
