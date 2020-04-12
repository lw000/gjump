package main

import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"github.com/lw000/gocommon/sys"
	log "github.com/sirupsen/logrus"
	"gjump/dao/database"
	"gjump/dao/service"
	"gjump/global"
	"os"
)

func main() {
	// 注册函数
	tysys.RegisterOnInterrupt(func(sign os.Signal) {
		log.WithField("sign", fmt.Sprintf("%v", sign)).Error("TWEB·服务退出")
	})

	// 加载配置文件
	if err := global.LoadGlobalConfig(); err != nil {
		log.Panic(err)
	}

	var err error

	// 连接数据库
	err = database.OpenMysql(*global.ProjectConfig.MysqlCfg)
	if err != nil {
		log.Panic(err)
		return
	}

	// 加载渠道数据
	{
		serve := service.CanalListDaoService{}
		if err = serve.Preload(); err != nil {
			log.Error(err)
		}
	}

	// 加载节点服务器地址数据
	{
		serve := service.ApiListDaoService{}
		if err = serve.Preload(); err != nil {
			log.Error(err)
		}
	}

	setupGin()
}
