package storeProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/storeProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上架商品
func ShelfProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShelfProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storeProduct.NewShelfProductLogic(r.Context(), svcCtx)
		err := l.ShelfProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
