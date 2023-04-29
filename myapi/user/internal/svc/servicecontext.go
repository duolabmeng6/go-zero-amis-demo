package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"myapi/model"
	"myapi/user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
