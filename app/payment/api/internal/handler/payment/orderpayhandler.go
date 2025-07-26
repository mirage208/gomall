package payment

import (
	"net/http"

	"github.com/mirage208/gomall/app/payment/api/internal/logic/payment"
	"github.com/mirage208/gomall/app/payment/api/internal/svc"
	"github.com/mirage208/gomall/app/payment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单付款
func OrderPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderPayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := payment.NewOrderPayLogic(r.Context(), svcCtx)
		resp, err := l.OrderPay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
