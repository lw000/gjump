package middleware

import (
	tyutils "github.com/lw000/gocommon/utils"
	"gjump/config"
	"gjump/dao/service"
	"gjump/models"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// TODO: 验证游客登录
func ValidateVisitorHandler(serveConf config.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("web:server:conf", serveConf)

		if serveConf.Config.VisitorLogin {
			_, err := c.Request.Cookie("clientuuid")
			if err == http.ErrNoCookie {
				registerIP := c.ClientIP()
				if registerIP == "" {
					log.Error("获取客户端IP地址错误")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "非法访问"})
					return
				}

				loginArgs := models.LoginArgs{}
				err = loginArgs.New(serveConf.Config.Agent, serveConf.Config.Canal, "0")
				if err != nil {
					log.Error(err.Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "创建游客账号信息错误"})
					return
				}

				cookieKey := tyutils.UUID()

				daoServe := service.CookieDaoService{}
				if err = daoServe.Save(cookieKey, loginArgs.URLString()); err != nil {
					log.Error(err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
					return
				}

				m := make(map[string]string)
				m["clientuuid"] = cookieKey
				m["visitorLogin"] = "1"
				for k, v := range m {
					ck := &http.Cookie{
						Name:     k,
						Value:    v,
						HttpOnly: false,
					}
					c.SetCookie(k, v, ck.MaxAge, ck.Path, ck.Domain, ck.Secure, ck.HttpOnly)
					c.Request.AddCookie(ck)
				}
				log.WithFields(log.Fields{"loginArgs": loginArgs.JSON()}).Info("游客登录生成参数")
				c.Set("Cookie", loginArgs)
			}
		}
		c.Next()
	}
}

func NoCacheHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Cache-Control", "no-cache")
		c.Next()
	}
}
