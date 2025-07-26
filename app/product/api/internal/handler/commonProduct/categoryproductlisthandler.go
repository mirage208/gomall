package commonProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/commonProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 商品分类列表
func CategoryProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryProductListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := commonProduct.NewCategoryProductListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryProductList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
