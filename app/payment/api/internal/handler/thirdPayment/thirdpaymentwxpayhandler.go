package thirdPayment

import (
	"net/http"

	"github.com/mirage208/gomall/app/payment/api/internal/logic/thirdPayment"
	"github.com/mirage208/gomall/app/payment/api/internal/svc"
	"github.com/mirage208/gomall/app/payment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// third payment：wechat pay
func ThirdPaymentwxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThirdPaymentWxPayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := thirdPayment.NewThirdPaymentwxPayLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPaymentwxPay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
