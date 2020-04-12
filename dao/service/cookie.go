package service

import "gjump/dao/ldb"

type CookieDaoService struct {
}

// Query ...
func (daoServe *CookieDaoService) Query(key string) (string, error) {
	serve := ldb.CookieCache{}
	return serve.Query(key)
}

// Save ...
func (daoServe *CookieDaoService) Save(key string, value string) error {
	serve := ldb.CookieCache{}
	return serve.Save(key, value)
}
