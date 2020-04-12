package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lw000/gocommon/utils"
	log "github.com/sirupsen/logrus"
	"gjump/config"
	"gjump/dao"
	"html/template"
	"net/http"
	"net/url"
)

// CustomerLoginHandler ...
func CustomerLoginHandler(c *gin.Context) {
	title := "大洋棋牌"
	query := c.Request.URL.RawQuery
	if query == "" {
		log.Error("登录参数为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	var param url.Values
	param, err := url.ParseQuery(query)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	if len(param) == 0 {
		log.Error("登录参数为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	account := param.Get("account")
	if account == "" {
		log.Error("账号为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	gameId := param.Get("gameId")
	if gameId == "" {
		log.Error("游戏ID为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	scanalId := param.Get("canalId")
	if scanalId == "" {
		log.Error("渠道为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	token := param.Get("token")
	if token == "" {
		log.Error("token为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	node := param.Get("node")
	if node == "" {
		log.Error("node为空")
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	canalId, err := tyutils.ToInt32(scanalId)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	vvv, exists := c.Get("web:server:conf")
	if !exists {
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}
	serveConf := vvv.(config.Server)

	platformId, err := tyutils.ToInt32(serveConf.Config.Platform)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
	}

	var gameUrl string
	gameUrl, _, err = dao.QueryServiceAddress(platformId, canalId)
	if err != nil {
		log.Error(fmt.Sprintf("渠道[%d], %s", canalId, err.Error()))
		c.HTML(http.StatusOK, "error.html", gin.H{"title": template.HTML(title), "msg": template.HTML("获取游戏启动参数失败")})
		return
	}

	gameURL := fmt.Sprintf("%s?%s&returnUrl=", gameUrl, query)

	log.WithFields(log.Fields{"canalId": canalId, "game": gameURL}).Info("H5·游戏地址")

	// c.Redirect(http.StatusMovedPermanently, gameURL)
	c.HTML(http.StatusOK, "index.html", gin.H{"title": template.HTML(title), "url": template.HTML(gameURL)})
}
