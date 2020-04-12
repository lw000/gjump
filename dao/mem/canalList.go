package mem

import (
	"fmt"
	"gjump/cache"
	"gjump/dao/table"
)

// 渠道缓存对象 ...
type CanalListCacheService struct {
	PId     int32
	CanalId int32
}

// key 缓存KEY
func (serve *CanalListCacheService) key() string {
	key := fmt.Sprintf("canalList:%d:%d", serve.PId, serve.CanalId)
	return key
}

// Exists ...
func (serve *CanalListCacheService) Exists() bool {
	exist := cache.MEMCacheService().Exist(serve.key())
	return exist
}

// Load 缓存中读取数据
func (serve *CanalListCacheService) Load() (table.TCanalList, error) {
	v, err := cache.MEMCacheService().Get(serve.key())
	if err != nil {
		return table.TCanalList{}, err
	}
	canalList := v.(table.TCanalList)
	return canalList, nil
}

// Save 保存渠道数据到Cache中
func (serve *CanalListCacheService) Save(canals table.TCanalList) error {
	if err := cache.MEMCacheService().Set(serve.key(), canals); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (serve *CanalListCacheService) Clear() bool {
	return cache.MEMCacheService().Remove(serve.key())
}
