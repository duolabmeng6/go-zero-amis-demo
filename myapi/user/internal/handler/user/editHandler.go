package user

import (
	"myapi/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/user/internal/logic/user"
	"myapi/user/internal/svc"
	"myapi/user/internal/types"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEditLogic(r.Context(), svcCtx)
		resp, err := l.Edit(&req)

		response.Response(w, resp, err)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
