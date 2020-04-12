package config

import (
	"encoding/json"
	tymysql "github.com/lw000/gocommon/db/mysql"
	"io/ioutil"
)

type Server struct {
	Listen      int64    `json:"listen"`
	Servername  []string `json:"servername"`
	Blacklist   []string `json:"blacklist"`
	Whitelist   []string `json:"whitelist"`
	Ssl         string   `json:"ssl"`
	SslCertfile string   `json:"ssl_certfile"`
	SslKeyfile  string   `json:"ssl_keyfile"`
	Config      struct {
		Agent        string `json:"agent"`
		Canal        string `json:"canal"`
		Platform     string `json:"platform"`
		VisitorLogin bool   `json:"visitorLogin"`
	} `json:"config"`
}

type Servers struct {
	Debug  int64
	Server []Server
}

// JSONConfig ...
type JSONConfig struct {
	MysqlCfg *tymysql.JsonConfig
	Servers  Servers
}

// NewJSONConfig ...
func NewJSONConfig() *JSONConfig {
	return &JSONConfig{
		MysqlCfg: &tymysql.JsonConfig{},
	}
}

// LoadJSONConfig ...
func LoadJSONConfig(file string) (*JSONConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var ccf CfgStruct
	if err = json.Unmarshal(data, &ccf); err != nil {
		return nil, err
	}

	cfg := NewJSONConfig()
	for _, serv := range ccf.Servers.Server {
		cfg.Servers.Server = append(cfg.Servers.Server, Server{
			serv.Listen,
			serv.Servername,
			serv.Blacklist,
			serv.Whitelist,
			serv.Ssl,
			serv.SslCertfile,
			serv.SslKeyfile,
			struct {
				Agent        string `json:"agent"`
				Canal        string `json:"canal"`
				Platform     string `json:"platform"`
				VisitorLogin bool   `json:"visitorLogin"`
			}{
				Agent:        serv.Config.Agent,
				Canal:        serv.Config.Canal,
				Platform:     serv.Config.Platform,
				VisitorLogin: serv.Config.VisitorLogin,
			},
		})
	}

	// 数据库配置
	cfg.MysqlCfg.Database = ccf.Mysql.Database
	cfg.MysqlCfg.Host = ccf.Mysql.Host
	cfg.MysqlCfg.MaxOdleConns = ccf.Mysql.MaxOdleConns
	cfg.MysqlCfg.MaxOpenConns = ccf.Mysql.MaxOpenConns
	cfg.MysqlCfg.Password = ccf.Mysql.Password
	cfg.MysqlCfg.Username = ccf.Mysql.Username

	return cfg, err
}
