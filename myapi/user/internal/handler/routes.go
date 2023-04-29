// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "myapi/user/internal/handler/user"
	"myapi/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/edit",
				Handler: user.EditHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
