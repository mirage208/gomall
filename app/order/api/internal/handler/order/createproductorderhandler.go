package order

import (
	"net/http"

	"github.com/mirage208/gomall/app/order/api/internal/logic/order"
	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建普通商品订单
func CreateProductOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProductOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewCreateProductOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateProductOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
