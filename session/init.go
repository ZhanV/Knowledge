package session

import "fmt"

var sessionMgr SessionMgr

func Init(provider string , addr string, options ...string) error {
	switch provider {
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	default:
		fmt.Errorf("not support provider : %s", provider)
	}
	return sessionMgr.Init(addr, options...)
}
