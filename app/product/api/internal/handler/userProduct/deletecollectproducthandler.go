package userProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/userProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除收藏商品
func DeleteCollectProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCollectProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userProduct.NewDeleteCollectProductLogic(r.Context(), svcCtx)
		err := l.DeleteCollectProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
