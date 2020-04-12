package table

// (apilist)节点服务器地址表
type TApiList struct {
	ApiListId int32  `json:"apiListId"` // id
	CanalId   int32  `json:"canalId"`   // 渠道id
	Domain    string `json:"domain"`    // 域名，带HTTP开头
	// Ip         string       `json:"ip"`         //
	// Remarks    string       `json:"remarks"`    //
	// Updatetime sql.NullTime `json:"updatetime"` // 修改时间
}
