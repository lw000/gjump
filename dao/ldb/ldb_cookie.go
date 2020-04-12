package ldb

import (
	log "github.com/sirupsen/logrus"
	"gjump/cache"
)

type CookieCache struct {
}

// Query ...
func (c *CookieCache) Query(key string) (string, error) {
	cookie, err := cache.LDBCacheService().Get(key)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return string(cookie), nil
}

// Save ...
func (c *CookieCache) Save(key string, cookie string) error {
	if err := cache.LDBCacheService().Set(key, []byte(cookie)); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (c *CookieCache) Clear() {

}
