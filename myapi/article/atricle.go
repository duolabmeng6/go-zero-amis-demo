package main

import (
	"flag"
	"fmt"
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
	}), rest.WithCors("*"))

	defer server.Stop()
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
