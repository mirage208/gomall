package order

import (
	"net/http"

	"github.com/mirage208/gomall/app/order/api/internal/logic/order"
	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单列表
func OrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderListLogic(r.Context(), svcCtx)
		resp, err := l.OrderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
