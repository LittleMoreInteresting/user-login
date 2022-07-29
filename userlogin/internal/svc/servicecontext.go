package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"user-login/userlogin/internal/config"
	"user-login/userlogin/internal/middleware"
	"user-login/userlogin/model/tags"
	"user-login/userlogin/model/user"
)

type ServiceContext struct {
	Config    config.Config
	Tagging   rest.Middleware
	Version   rest.Middleware
	UserModel user.UserModel
	TagsModel tags.BlogTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(conn, c.CacheRedis),
		TagsModel: tags.NewBlogTagModel(conn, c.CacheRedis),
		Tagging:   middleware.NewTaggingMiddleware().Handle,
		Version:   middleware.NewVersionMiddleware().Handle,
	}
}
