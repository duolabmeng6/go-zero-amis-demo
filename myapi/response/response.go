package response

import (
	"myapi/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err interface{}) {

	var body Body
	switch e := err.(type) {
	case *errorx.CodeError:
		body.Code = e.Data().Code
		body.Msg = e.Data().Msg
	default:
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
