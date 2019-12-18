package account

import "github.com/zhanv/knowledge/session"

func InitSession(provider string , addr string, options ...string) error {
	session.Init(provider,addr,options...)
}
