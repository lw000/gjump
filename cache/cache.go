package cache

import (
	tycache "github.com/lw000/gocommon/cache"
	"sync"
)

var (
	memCache     tycache.MemCache
	memCacheOnce sync.Once
)

// MEMCacheService ...
func MEMCacheService() tycache.MemCache {
	memCacheOnce.Do(func() {
		memCache = tycache.NewLRU(1024 * 100)
	})
	return memCache
}
