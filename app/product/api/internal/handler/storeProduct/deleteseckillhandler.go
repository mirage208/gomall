package storeProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/storeProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/mirage208/gomall/app/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除秒杀商品
func DeleteSeckillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSeckillProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storeProduct.NewDeleteSeckillLogic(r.Context(), svcCtx)
		err := l.DeleteSeckill(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
