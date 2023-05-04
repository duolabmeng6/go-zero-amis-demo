package tool

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
)

// StaticFileHandler 静态文件服务器
func StaticFileHandler(server *rest.Server, pathPrefix string, fileDir string) {
	//Parameters:
	//server 服务器对象
	// pathPrefix: 指定静态文件服务器的路由前缀，即请求路径中的前缀部分，用于标识不同的静态文件。
	// fileDir: 指定静态文件服务器的文件目录，即存放静态文件的地方。

	// 目录的层级 10层基本够用了吧
	dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8", ":9", ":10"}
	for i := 1; i < len(dirlevel); i++ {
		// 注册这种路由 /:1/:2/:3/:4/:5/:6/:7"
		pathx := "/" + strings.Join(dirlevel[:i], "/")
		//logx.Info(pathx)
		//最后生成 /asset
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    pathx,
				Handler: StaticServerHandler(pathPrefix, fileDir),
			})
	}
}

func StaticServerHandler(pathPrefix, dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler := http.StripPrefix(pathPrefix, http.FileServer(http.Dir(dir)))
		handler.ServeHTTP(w, r)
	}
}
