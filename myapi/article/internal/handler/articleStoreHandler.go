package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/article/internal/logic"
	"myapi/article/internal/svc"
	"myapi/article/internal/types"
)

func ArticleStoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleStoreRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleStoreLogic(r.Context(), svcCtx)
		resp, err := l.ArticleStore(&req)

		//response.Response(w, resp, err)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
