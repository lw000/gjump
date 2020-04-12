package api

import (
	"gjump/config"
	"gjump/routers/api/handler"

	"github.com/gin-gonic/gin"
)

// IndexHandler ...
func IndexHandler(c *gin.Context) {
	vvv, exists := c.Get("web:server:conf")
	if !exists {
		return
	}
	webServeConf := vvv.(config.Server)

	if webServeConf.Config.VisitorLogin {
		handler.VisitorLoginHandler(c)
	} else {
		handler.CustomerLoginHandler(c)
	}
}
