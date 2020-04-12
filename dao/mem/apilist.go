package mem

import (
	"fmt"
	"gjump/cache"
	"gjump/dao/table"
)

// 渠道缓存对象 ...
type ApiListCacheService struct {
	CanalId int32
}

// key 缓存KEY
func (serve *ApiListCacheService) key() string {
	key := fmt.Sprintf("apiList:%d", serve.CanalId)
	return key
}

// Exists ...
func (serve *ApiListCacheService) Exists() bool {
	return cache.MEMCacheService().Exist(serve.key())
}

// Load 缓存中读取数据
func (serve *ApiListCacheService) Load() ([]table.TApiList, error) {
	v, err := cache.MEMCacheService().Get(serve.key())
	if err != nil {
		return nil, err
	}
	apis := v.([]table.TApiList)
	return apis, nil
}

// Save 保存渠道数据到Cache中
func (serve *ApiListCacheService) Save(apis ...table.TApiList) error {
	if err := cache.MEMCacheService().Set(serve.key(), apis); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (serve *ApiListCacheService) Clear() bool {
	return cache.MEMCacheService().Remove(serve.key())
}
