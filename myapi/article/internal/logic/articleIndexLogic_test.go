package logic

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"myapi/article/internal/config"
	"myapi/article/internal/svc"
	"myapi/article/internal/types"
	"testing"
)

var configFile = flag.String("f", "../../etc/atricle-api.yaml", "the config file")

func TestIndex(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	logic := NewArticleIndexLogic(context.Background(), svcCtx)
	resp, err := logic.ArticleIndex(&types.ArticleIndexRequest{
		Page:     1,
		PerPage:  10,
		Keywords: "",
		OrderBy:  "id",
		OrderDir: "desc",
	})
	//优雅的输出结构体
	t.Logf("%+v", resp)
	t.Logf("%+v", err)

}
