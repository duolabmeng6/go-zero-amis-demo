package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"myapi/article/internal/config"
	"myapi/article/internal/middleware"
	"myapi/model"
)

type ServiceContext struct {
	Config         config.Config
	ArticleModel   model.ArticlesModel
	CORSMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ArticleModel:   model.NewArticlesModel(sqlx.NewMysql(c.DB.DataSource)),
		CORSMiddleware: middleware.NewCORSMiddleware().Handle,
	}
}
