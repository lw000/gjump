package cache

import (
	"github.com/lw000/gocommon/cache"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	ldbCache     tycache.LeveldbCache
	ldbCacheOnce sync.Once
)

func newLeveldbCache(path string) tycache.LeveldbCache {
	c, err := tycache.NewLDB(path)
	if err != nil {
		log.Error(err)
		return nil
	}
	return c
}

// LDBCacheService 创建新的DB实例
func LDBCacheService() tycache.LeveldbCache {
	ldbCacheOnce.Do(func() {
		ldbCache = newLeveldbCache("./data/ldb/")
	})
	return ldbCache
}
