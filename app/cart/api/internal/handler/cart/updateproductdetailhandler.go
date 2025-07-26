package cart

import (
	"net/http"

	"github.com/mirage208/gomall/app/cart/api/internal/logic/cart"
	"github.com/mirage208/gomall/app/cart/api/internal/svc"
	"github.com/mirage208/gomall/app/cart/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改购物车商品的信息
func UpdateProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewUpdateProductDetailLogic(r.Context(), svcCtx)
		err := l.UpdateProductDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
