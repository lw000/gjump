package routers

import (
	"gjump/config"
	"gjump/middleware"
	"gjump/routers/api"
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

type Controller1 gin.HandlerFunc

func (control Controller1) Run(c *gin.HandlerFunc) {

}

type Controller struct {
}

type LoginController struct {
	Controller
}

func (control *LoginController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "1")
	}
}

// RegisterService ...
func RegisterService(engine *gin.Engine, serv config.Server) {
	lmt := tollbooth.NewLimiter(200, nil)
	lmt.SetMethods([]string{"GET", "POST"})
	lmt.SetMessage("too many requests")

	engine.GET("/", tollbooth_gin.LimitHandler(lmt), middleware.ValidateVisitorHandler(serv), api.IndexHandler)

	engine.GET("/test", (&LoginController{}).Index())
}
