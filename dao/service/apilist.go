package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gjump/dao/database"
	"gjump/dao/mem"
	"gjump/dao/table"
)

type ApiListDaoService struct {
}

// 加载渠道数据
func (daoServe *ApiListDaoService) Preload() error {
	apis, err := daoServe.Select()
	if err != nil {
		log.Error(err)
		return err
	}

	for canalId, c := range apis {
		cacheServe := mem.ApiListCacheService{CanalId: canalId}
		if err = cacheServe.Save(c...); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

// 查询渠道信息
func (daoServe *ApiListDaoService) Query(canalId int32, ids ...string) ([]table.TApiList, error) {
	// 1. 查询缓存数据
	cacheServe := mem.ApiListCacheService{CanalId: canalId}
	apis, err := cacheServe.Load()
	if err != nil {
		log.WithFields(log.Fields{"canalId": canalId}).Error(err)
	}

	if len(apis) == 0 {
		// 2. 查询数据库
		apis, err = daoServe.selectWithCanalId(canalId)
		if err != nil {
			log.WithFields(log.Fields{"canalId": canalId}).Error(err)
			return nil, err
		}

		// 3. 更新缓存
		if err = cacheServe.Save(apis...); err != nil {
			log.WithFields(log.Fields{"canalId": canalId}).Error(err)
		}
	}

	return daoServe.filterApiList(apis, ids...)
}

func (daoServe *ApiListDaoService) selectWithCanalId(canalId int32) ([]table.TApiList, error) {
	query := `SELECT apiListId, canalId, domain
					FROM apiList
						WHERE canalId =?;`
	rows, err := database.GetMysql(database.DB_GAMEDATA).DB().Query(query, canalId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	canals := make([]table.TApiList, 0, 8)
	for rows.Next() {
		var apis table.TApiList
		err = rows.Scan(
			&apis.ApiListId,
			&apis.CanalId,
			&apis.Domain,
		)
		if err == nil {
			canals = append(canals, apis)
		} else {
			log.Error(err)
		}
	}
	return canals, nil
}

func (daoServe *ApiListDaoService) Select() (map[int32][]table.TApiList, error) {
	query := `SELECT apiListId, canalId, domain FROM apiList;`
	rows, err := database.GetMysql(database.DB_GAMEDATA).DB().Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	m := make(map[int32][]table.TApiList, 64)
	for rows.Next() {
		var lst table.TApiList
		err = rows.Scan(
			&lst.ApiListId,
			&lst.CanalId,
			&lst.Domain,
		)
		if err == nil {
			var exists bool
			_, exists = m[lst.CanalId]
			if !exists {
				m[lst.CanalId] = make([]table.TApiList, 0, 8)
			}
			m[lst.CanalId] = append(m[lst.CanalId], lst)
		} else {
			log.Error(err)
		}
	}
	return m, nil
}

// filterApiList
func (daoServe *ApiListDaoService) filterApiList(apis []table.TApiList, ids ...string) ([]table.TApiList, error) {
	for _, id := range ids {
		for i, u := range apis {
			if fmt.Sprintf("%d", u.ApiListId) != id {
				apis = append(apis[0:i], apis[i+1:]...)
			}
		}
	}
	return apis, nil
}
