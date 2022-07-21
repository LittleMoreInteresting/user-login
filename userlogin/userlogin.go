package main

import (
	"flag"
	"fmt"

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

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
