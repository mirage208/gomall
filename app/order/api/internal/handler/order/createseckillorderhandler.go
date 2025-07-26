package order

import (
	"net/http"

	"github.com/mirage208/gomall/app/order/api/internal/logic/order"
	"github.com/mirage208/gomall/app/order/api/internal/svc"
	"github.com/mirage208/gomall/app/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建秒杀商品订单
func CreateSeckillOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSeckillOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewCreateSeckillOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateSeckillOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
