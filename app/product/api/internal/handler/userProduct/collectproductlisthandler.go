package userProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/userProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 收藏商品列表
func CollectProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userProduct.NewCollectProductListLogic(r.Context(), svcCtx)
		resp, err := l.CollectProductList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
