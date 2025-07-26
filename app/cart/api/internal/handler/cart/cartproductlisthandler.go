package cart

import (
	"net/http"

	"github.com/mirage208/gomall/app/cart/api/internal/logic/cart"
	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 购物车商品列表
func CartProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewCartProductListLogic(r.Context(), svcCtx)
		resp, err := l.CartProductList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
