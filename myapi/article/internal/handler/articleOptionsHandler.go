package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/article/internal/logic"
	"myapi/article/internal/svc"
)

func ArticleOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewArticleOptionsLogic(r.Context(), svcCtx)
		err := l.ArticleOptions()

		//response.Response(w, resp, err)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
