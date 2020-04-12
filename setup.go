package main

import (
	"fmt"
	"github.com/lw000/gocommon/app/gin"
	"github.com/lw000/gocommon/web/gin/middleware"
	"gjump/config"
	"gjump/global"
	"gjump/middleware"
	"gjump/routers"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func ginHandler(server config.Server) http.Handler {
	app := tygin.NewApplication(global.ProjectConfig.Servers.Debug)
	app.Init()
	// 跨域处理
	app.Engine().Use(tymiddleware.CorsHandler(nil))
	// 主机域名绑定
	app.Engine().Use(tymiddleware.HostBindingHandler(server.Servername))
	// 禁止缓存
	app.Engine().Use(middleware.NoCacheHandler())

	app.Engine().Static("/public", "./public")
	app.Engine().StaticFile("/favicon.ico", "./public/favicon.ico")

	app.Engine().LoadHTMLGlob("templates/*")

	routers.RegisterService(app.Engine(), server)

	return app.Engine()
}

func setupGin() {
	var g errgroup.Group

	for _, serv := range global.ProjectConfig.Servers.Server {
		servCopy := serv
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", servCopy.Listen),
			Handler: ginHandler(servCopy),
		}

		if strings.ToLower(servCopy.Ssl) == "on" {
			g.Go(func() error {
				log.Info("Listening and serving HTTPS on ", server.Addr)

				if err := server.ListenAndServeTLS(servCopy.SslCertfile, servCopy.SslKeyfile); err != nil {
					log.Error(err)
					return err
				}
				return nil
			})
		} else {
			g.Go(func() error {
				log.Info("Listening and serving HTTP on ", server.Addr)

				if err := server.ListenAndServe(); err != nil {
					log.Error(err)
					return err
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		log.Panic(err)
	}
}
