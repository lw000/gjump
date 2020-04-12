package service

import (
	"database/sql"
	"gjump/dao/database"
	"gjump/dao/mem"
	"gjump/dao/table"
	"gjump/errors"

	log "github.com/sirupsen/logrus"
)

type CanalListDaoService struct {
}

// 加载渠道数据
func (daoServe *CanalListDaoService) Preload() error {
	canals, err := daoServe.Select()
	if err != nil {
		log.Error(err)
		return err
	}

	for _, canal := range canals {
		cacheServe := mem.CanalListCacheService{PId: canal.PlatformId, CanalId: canal.CanalId}
		if err = cacheServe.Save(canal); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

// 查询渠道信息
func (daoServe *CanalListDaoService) Query(platformId int32, canalId int32) (table.TCanalList, error) {
	// 1. 查询缓存数据
	cacheServe := mem.CanalListCacheService{PId: platformId, CanalId: canalId}
	canal, err := cacheServe.Load()
	if err != nil {
		log.WithFields(log.Fields{"platformId": platformId, "canalId": canalId}).Error(err)
	}

	if canal.CanalId == 0 {
		// 2. 查询数据库
		canal, err = daoServe.SelectWith(platformId, canalId)
		if err != nil {
			log.WithFields(log.Fields{"platformId": platformId, "canalId": canalId}).Error(err)
			return table.TCanalList{}, err
		}

		// 3. 更新缓存
		if err = cacheServe.Save(canal); err != nil {
			log.WithFields(log.Fields{"platformId": platformId, "canalId": canalId}).Error(err)
		}
	}

	return canal, nil
}

// SelectWith
func (daoServe *CanalListDaoService) SelectWith(platformId int32, canalId int32) (table.TCanalList, error) {
	query := `SELECT
				canalId,
				canalName,
       			packageType,
				h5GameUrl,
				apiUrl,
				platformId
			FROM
				canalList
			WHERE
				platformId =? AND canalId =? AND packageType=2;`
	row := database.GetMysql(database.DB_GAMEDATA).DB().QueryRow(query, platformId, canalId)

	var canal table.TCanalList
	err := row.Scan(
		&canal.CanalId,
		&canal.CanalName,
		&canal.PackageType,
		&canal.H5GameUrl,
		&canal.ApiUrl,
		&canal.PlatformId,
	)
	if err != nil && err != sql.ErrNoRows {
		log.WithFields(log.Fields{"platformId": platformId, "canalId": canalId}).Error(err)
		return table.TCanalList{}, err
	}

	if err == sql.ErrNoRows {
		return table.TCanalList{}, errors.New(0, "未查询到渠道", "")
	}

	return canal, nil
}

func (daoServe *CanalListDaoService) Select() ([]table.TCanalList, error) {
	query := `SELECT
				canalId,
				canalName,
       			packageType,
				h5GameUrl,
				apiUrl,
				platformId
			FROM
				canalList
			WHERE packageType=2;`
	rows, err := database.GetMysql(database.DB_GAMEDATA).DB().Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	canals := make([]table.TCanalList, 0, 128)
	for rows.Next() {
		var canal table.TCanalList
		err = rows.Scan(
			&canal.CanalId,
			&canal.CanalName,
			&canal.PackageType,
			&canal.H5GameUrl,
			&canal.ApiUrl,
			&canal.PlatformId,
		)
		if err == nil {
			canals = append(canals, canal)
		} else {
			log.Error(err)
		}
	}
	return canals, nil
}
