package commonProduct

import (
	"net/http"

	"github.com/mirage208/gomall/app/product/api/internal/logic/commonProduct"
	"github.com/mirage208/gomall/app/product/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 首页分类名称列表
func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := commonProduct.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
