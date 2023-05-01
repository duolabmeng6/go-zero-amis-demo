package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/article/internal/config"
	"myapi/article/internal/handler"
	"myapi/article/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/atricle-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println("UnauthorizedCallback")
		httpx.Error(w, errors.New("UnauthorizedCallback"))
	}), rest.WithCors("*"))

	defer server.Stop()
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
