package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user-login/userlogin/internal/config"
	"user-login/userlogin/model/user"
)

type ServiceContext struct {
	Config config.Config

	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(conn, c.CacheRedis),
	}
}
