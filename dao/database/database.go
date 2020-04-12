package database

import (
	tymysql "github.com/lw000/gocommon/db/mysql"
	"sync"

	log "github.com/sirupsen/logrus"
)

const (
	DB_GAMEDATA = "gamedata"
)

var (
	m        sync.RWMutex
	dbServer map[string]*tymysql.Mysql
)

func init() {
	dbServer = make(map[string]*tymysql.Mysql)
}

// 打开数据库
func OpenMysql(configs ...tymysql.JsonConfig) error {
	var g sync.WaitGroup
	for _, cfg := range configs {
		cfg := cfg
		g.Add(1)
		go func() {
			defer func() {
				g.Done()
			}()
			myDb := &tymysql.Mysql{}
			if err := myDb.OpenWithJsonConfig(&cfg); err != nil {
				return
			}
			m.Lock()
			dbServer[cfg.Database] = myDb
			m.Unlock()
			log.WithFields(log.Fields{"database": cfg.Database}).Info("数据库连接成功")
		}()
	}
	g.Wait()
	return nil
}

func CloseMysql() {
	for _, myDb := range dbServer {
		_ = myDb.Close()
	}
}

// 获取数据库连接
func GetMysql(database string) *tymysql.Mysql {
	m.RLock()
	myDb, exists := dbServer[database]
	if exists {
		m.RUnlock()
		return myDb
	}
	m.RUnlock()
	return nil
}
