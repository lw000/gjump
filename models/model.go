package models

import (
	"encoding/json"
	"fmt"
	"github.com/lw000/gocommon/auth"
	"github.com/lw000/gocommon/auth/tymd5"
	"github.com/lw000/gocommon/utils"
	"gjump/errors"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// LoginArgs ...
type LoginArgs struct {
	Account string `json:"account"`
	Agent   string `json:"agent"`
	CanalId string `json:"canalId"`
	GameId  string `json:"gameId"`
	Token   string `json:"token"`
	Node    string `json:"node"`
	Visitor string `json:"visitor"`
}

// NewArgs ...
func (l *LoginArgs) New(agent, canalId, gameId string) error {
	var err error
	l.Account, err = tymd5.MD5([]byte(tyutils.UUID()))
	if err != nil {
		log.Error("生成游客账号失败")
		return errors.New(0, "生成游客账号失败", "")
	}

	l.Token, err = tymd5.MD5([]byte(tyutils.UUID()))
	if err != nil {
		log.Error("生成token失败")
		return errors.New(0, "生成token失败", "")
	}

	l.Agent = agent
	l.CanalId = canalId
	l.GameId = gameId
	l.Visitor = "1"

	return nil
}

// NewWithJSON ...
func (l *LoginArgs) NewWithJSON(data []byte) error {
	if err := json.Unmarshal(data, l); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// NewWithURLQuery ...
func (l *LoginArgs) NewWithURLQuery(query string) error {
	u, err := url.ParseQuery(query)
	if err != nil {
		log.Error(err)
		return err
	}
	l.Account = u.Get("account")
	l.Agent = u.Get("agent")
	l.CanalId = u.Get("canalId")
	l.GameId = u.Get("gameId")
	l.Token = u.Get("token")
	l.Node = u.Get("node")
	l.Visitor = u.Get("visitor")

	return nil
}

func (l *LoginArgs) String() string {
	return fmt.Sprintf("{account:%s agent:%s canalId:%s gameId:%s token:%s}", l.Account, l.Agent, l.CanalId, l.GameId, l.Token)
}

// JSON ...
func (l *LoginArgs) JSON() string {
	data, err := json.Marshal(l)
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(data)
}

// URLString ...
func (l *LoginArgs) URLString() string {
	vls := url.Values{}
	vls.Add("account", l.Account)
	vls.Add("agent", l.Agent)
	vls.Add("canalId", l.CanalId)
	vls.Add("gameId", l.GameId)
	vls.Add("token", l.Token)
	vls.Add("node", l.Node)
	vls.Add("visitor", l.Visitor)
	return vls.Encode()
}

// EncryptURLString ...
func (l *LoginArgs) EncryptURLString() string {
	s := l.URLString()
	s = tyutils.Reverse(tyauth.Hex([]byte(s)))
	return s
}
