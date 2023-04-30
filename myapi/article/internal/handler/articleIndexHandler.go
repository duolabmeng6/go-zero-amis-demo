package handler

import (
	"myapi/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/article/internal/logic"
	"myapi/article/internal/svc"
	"myapi/article/internal/types"
)

func ArticleIndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleIndexRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleIndexLogic(r.Context(), svcCtx)
		resp, err := l.ArticleIndex(&req)

		response.Response(w, resp, err)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
