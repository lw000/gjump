package config

type CfgStruct struct {
	Mysql struct {
		MaxOdleConns int64  `json:"MaxOdleConns"`
		MaxOpenConns int64  `json:"MaxOpenConns"`
		Database     string `json:"database"`
		Host         string `json:"host"`
		Password     string `json:"password"`
		Username     string `json:"username"`
	} `json:"mysql"`
	Servers struct {
		Debug  int64 `json:"debug"`
		Server []struct {
			Blacklist []string `json:"blacklist"`
			Config    struct {
				Agent        string `json:"agent"`
				Canal        string `json:"canal"`
				Platform     string `json:"platform"`
				VisitorLogin bool   `json:"visitorLogin"`
			} `json:"config"`
			Listen      int64    `json:"listen"`
			Servername  []string `json:"servername"`
			Ssl         string   `json:"ssl"`
			SslCertfile string   `json:"ssl_certfile"`
			SslKeyfile  string   `json:"ssl_keyfile"`
			Whitelist   []string `json:"whitelist"`
		} `json:"server"`
	} `json:"servers"`
}
