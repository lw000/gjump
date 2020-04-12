package table

// TCanalList 渠道数据结构
type TCanalList struct {
	CanalId   int32  `json:"canalId"`   // 渠道Id
	CanalName string `json:"canalName"` // 渠道名称
	// CanalNotice                string `json:"canalNotice"`                // 渠道公告
	// InitScore                  int32  `json:"initScore"`                  // 初始化金币
	// CanalStatus  int32  `json:"canalStatus"`  // 渠道状态，0维护，1正常，默认为1
	// CanalStopMsg string `json:"canalStopMsg"` // 渠道维护提示信息
	// CanalRate                  int32  `json:"canalRate"`                  // 渠道显示比例，要么1，要么10000
	// CanalType                  int32  `json:"canalType"`                  // 渠道类型。0-系统；1-第三方
	PackageType int32 `json:"packageType"` // 包类型：0-安卓；1-苹果；2-H5;3-苹果上架
	// AllowReg                   int32  `json:"allowReg"`                   // 是否允许注册  0不允许，1允许 默认为1
	// RegMsgCheck                int32  `json:"regMsgCheck"`                // 注册账号短信验证 0，不验证 1 验证 默认为1
	// AllowGuest                 int32  `json:"allowGuest"`                 // 允许游客,0不允许 1允许，默认为1
	// AllowWeiXin                int32  `json:"allowWeiXin"`                // 允许微信登陆0不允许，1允许，默认为1
	// EmuStrategy                int32  `json:"emuStrategy"`                // 模拟器策略 0禁止注册赠送金币 1禁止在模拟器进入游戏 2正常使用
	// CanalGiveTemplateId        int32  `json:"canalGiveTemplateId"`        // 赠送模板
	// CanalRewardTemplateId      int32  `json:"canalRewardTemplateId"`      // 兑奖模板
	// CanalShowControlTemplateId int32  `json:"canalShowControlTemplateId"` // 大厅按钮显示模板
	// CanalDiscountTemplateId    int32  `json:"canalDiscountTemplateId"`    // 扣量模板
	H5GameUrl string `json:"h5GameUrl"` // H5游戏的地址,JSON格式保存
	// H5UserPrefix string `json:"h5UserPrefix"` // h5代理用户前缀
	// ApkUrl                     string `json:"apkUrl"`                     // 安卓下载地址
	// IosUrl                     string `json:"iosUrl"`                     // ios下载地址
	// ForceApkUrl                string `json:"forceApkUrl"`                // 强制安卓下载地址
	// ForceIosUrl                string `json:"forceIosUrl"`                // 强制苹果下载地址
	// ForceUpgradeNotice         string `json:"forceUpgradeNotice"`         // 强更提示信息
	// HomeUrl                    string `json:"homeUrl"`                    // 官网
	// IosOpenUrl                 string `json:"iosOpenUrl"`                 // 推广页面
	// ServiceUrl                 string `json:"serviceUrl"`                 // 客服地址，如果不为空，客户端弹出这个网址
	// SpreadUrl                  string `json:"spreadUrl"`                  // 推广地址
	// ServiceNumberId            int32  `json:"serviceNumberId"`            // 客服编号Id
	ApiUrl string `json:"apiUrl"` // JOSN形式的API获取数据地址
	// WxAppId     string `json:"wxAppId"`     // 微信登陆用到的APPID
	// WxAppSecret string `json:"wxAppSecret"` // 微信登陆用到的APPSECRET
	// CustomAttr                 string `json:"customAttr"`                 // 自定义属性JSON
	PlatformId int32 `json:"platformId"` // 平台ID
}
