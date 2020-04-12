package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gjump/dao/service"
	"gjump/dao/table"
	"golang.org/x/sync/singleflight"
	"math/rand"
)

var (
	singleFlightCanalList singleflight.Group
	singleFlightApiList   singleflight.Group
)

// TODO: 查询渠道数据
func queryCanalList(platformId int32, canalId int32) (table.TCanalList, error) {
	v, err, sh := singleFlightCanalList.Do(fmt.Sprintf("canalList_%d_%d", platformId, canalId), func() (i interface{}, e error) {
		serve := service.CanalListDaoService{}
		return serve.Query(platformId, canalId)
	})

	if err != nil {
		return table.TCanalList{}, err
	}

	if sh {
	}

	return v.(table.TCanalList), nil
}

// TODO: 查询节点服务器地址数据
func queryApiList(canalId int32, ids ...string) ([]table.TApiList, error) {
	v, err, sh := singleFlightApiList.Do(fmt.Sprintf("apilist_%d", canalId), func() (i interface{}, e error) {
		serve := service.ApiListDaoService{}
		return serve.Query(canalId, ids...)
	})

	if err != nil {
		return nil, err
	}

	if sh {
	}

	return v.([]table.TApiList), nil
}

// TODO: 查询节点服务地址
func QueryServiceAddress(platformId, canalId int32) (gameUrl string, nodeUrl string, err error) {
	// 1. 查询渠道信息
	var canal table.TCanalList
	canal, err = queryCanalList(platformId, canalId)
	if err != nil {
		return "", "", err
	}

	if canal.CanalId == 0 {
		err = errors.New("未查询到渠道")
		return "", "", err
	}

	// H5游戏地址
	var gameUrls []string
	err = json.Unmarshal([]byte(canal.H5GameUrl), &gameUrls)
	if err != nil {
		return "", "", err
	}

	if len(gameUrls) == 0 {
		err = errors.New("渠道未配置游戏地址")
		return "", "", err
	}

	// 2. 查询节点服务器地址
	var apis []table.TApiList
	apis, err = queryApiList(canalId, canal.ApiUrl)
	if err != nil {
		return "", "", err
	}
	if len(apis) == 0 {
		err = errors.New("未查询到节点服务器地址")
		return "", "", err
	}

	// 随机分配, 网关节点地址
	nodeUrlIndex := rand.Intn(len(apis))
	nodeUrl = apis[nodeUrlIndex].Domain

	// 随机分配, H5游戏地址
	gameUrlIndex := rand.Intn(len(gameUrls))
	gameUrl = gameUrls[gameUrlIndex]

	return
}
