package handler

import (
	"fmt"
	"github.com/lw000/gocommon/utils"
	"gjump/config"
	"gjump/dao"
	"gjump/dao/service"
	"gjump/models"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// VisitorLoginHandler ...
func VisitorLoginHandler(c *gin.Context) {
	var (
		err       error
		cookie    *http.Cookie
		loginArgs models.LoginArgs
	)

	title := "游客模式·" + "大洋棋牌"

	// 启用游客登录
	v, exists := c.Get("Cookie")
	if exists {
		loginArgs = v.(models.LoginArgs)
	} else {
		cookie, err = c.Request.Cookie("clientuuid")
		if err != nil {
			errText := "获取cookie失败"
			log.Error(errText)
			c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML(errText)})
			return
		}

		if cookie.Value == "" {
			errText := "获取cookie失败"
			log.Error(errText)
			c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML(errText)})
			return
		}

		var query string
		serve := service.CookieDaoService{}
		query, err = serve.Query(cookie.Value)
		if err != nil {
			log.Error(err.Error())
			c.HTML(http.StatusOK, "clear_cookie.html", gin.H{"title": template.HTML(title), "msg": template.HTML("账号过期，点击这里重试")})
			return
		}

		loginArgs = models.LoginArgs{}
		if err = loginArgs.NewWithURLQuery(query); err != nil {
			log.Error(err.Error())
			c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
			return
		}
	}

	// 跳转自定义游戏
	gameId := c.Query("gameId")
	if gameId != "" {
		loginArgs.GameId = gameId
	}

	var vvv interface{}
	vvv, exists = c.Get("web:server:conf")
	if !exists {
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}
	serveConf := vvv.(config.Server)

	var canalId int32
	canalId, err = tyutils.ToInt32(loginArgs.CanalId)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	var platformId int32
	platformId, err = tyutils.ToInt32(serveConf.Config.Platform)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
	}

	var gameUrl, nodeUrl string
	gameUrl, nodeUrl, err = dao.QueryServiceAddress(platformId, canalId)
	if err != nil {
		log.Error(fmt.Sprintf("渠道[%d], %s", canalId, err.Error()))
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	loginArgs.Node = nodeUrl

	gameURL := fmt.Sprintf("%s?param=%s&t=%d", gameUrl, loginArgs.EncryptURLString(), time.Now().UnixNano())

	log.WithFields(log.Fields{"canalId": canalId, "game": gameURL}).Info("H5·游戏地址")

	// c.Redirect(http.StatusMovedPermanently, gameURL)

	c.HTML(http.StatusOK, "index.html", gin.H{"title": template.HTML(title), "url": template.HTML(gameURL)})
}
