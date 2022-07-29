package main

import (
	"flag"
	"fmt"
	"net/http"

	"user-login/userlogin/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"user-login/userlogin/internal/config"
	"user-login/userlogin/internal/handler"
	"user-login/userlogin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userlogin-api.yaml", "the config file")

//go run userlogin.go -f etc/userlogin-api.yaml
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.LogConf)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			logx.Error(e)
			err := errorx.NewCodeError(http.StatusInternalServerError, e.Error()).(*errorx.CodeError)
			return http.StatusOK, err.Data()
		}
	})

	unauthorized := rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.WriteJson(w, http.StatusOK, errorx.NewCodeError(http.StatusUnauthorized, err.Error()))
	})
	server := rest.MustNewServer(c.RestConf, unauthorized)

	defer server.Stop()
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			logx.Info("global middleware")
			next(writer, request)
		}
	})
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
