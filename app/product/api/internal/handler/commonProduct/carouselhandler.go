package commonProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/commonProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 首页商品轮播图
func CarouselHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := commonProduct.NewCarouselLogic(r.Context(), svcCtx)
		resp, err := l.Carousel()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
