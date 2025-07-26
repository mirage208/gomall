package commonProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/commonProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 首页商品推荐列表
func RecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := commonProduct.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
